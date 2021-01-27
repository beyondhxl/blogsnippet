package sl

import (
	"math/rand"
)

const (
	SkipListMaxLevel = 32

	SkipListP = 0.25
)

type IElement interface {
	// UniqueId() return a uint64 is used for insertion/deletion/find. It needs to establish an order over all elements
	UniqueId() uint64

	// A string representation of the element. Can be used for pretty-printing the list. Otherwise just return an empty string.
	String() string
}

type IKey interface {
	Less() bool
}

//跳表层（同一层到下个结点的距离）
type SkipListLevel struct {
	forward *SkipListNode //后继
	span    uint64        //跨度
}

type SkipListNode struct {
	objID    int64   //唯一id
	score    float64 //分数
	backward *SkipListNode
	level    []*SkipListLevel
}

type Obj struct {
	key        int64
	attachment interface{}
	score      float64
}

type SkipList struct {
	header *SkipListNode
	tail   *SkipListNode
	length int64
	level  int16
}

type SkipListRank struct {
	dict map[int64]*Obj
	sl   *SkipList
}

type rangespec struct {
	min   float64
	max   float64
	minex int32
	maxex int32
}

type lexrangespec struct {
	minKey int64
	maxKey int64
	minex  int
	maxex  int
}

func NewSkipListNode(level int16, score float64, id int64) *SkipListNode {
	sn := &SkipListNode{
		score: score,
		objID: id,
		level: make([]*SkipListLevel, level),
	}

	for i := range sn.level {
		sn.level[i] = new(SkipListLevel)
	}

	return sn
}

func NewSkipList() *SkipList {
	return &SkipList{
		level:  1,
		header: NewSkipListNode(SkipListMaxLevel, 0, 0),
	}
}

func randomLevel() int16 {
	level := int16(1)

	for float32(rand.Int31()&0xFFFF) < (SkipListP * 0xFFFF) {
		level++
	}

	if level < SkipListMaxLevel {
		return level
	}

	return SkipListMaxLevel
}

func (this *SkipList) Insert(score float64, id int64) *SkipListNode {
	update := make([]*SkipListNode, SkipListMaxLevel)
	rank := make([]uint64, SkipListMaxLevel)
	node := this.header

	for i := this.level - 1; i >= 0; i-- {
		/* store rank that is crossed to reach the insert position */
		if i == this.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		if node.level[i] != nil {
			for node.level[i].forward != nil &&
				(node.level[i].forward.score < score ||
					(node.level[i].forward.score == score && node.level[i].forward.objID > id)) {
				rank[i] += node.level[i].span
				node = node.level[i].forward
			}
		}

		update[i] = node
	}

	level := randomLevel()
	if level > this.level {
		for i := this.level; i < level; i++ {
			rank[i] = 0
			update[i] = this.header
			update[i].level[i].span = uint64(this.length)
		}
		this.level = level
	}

	node = NewSkipListNode(level, score, id)
	for i := int16(0); i < level; i++ {
		//链表的插入
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node

		/* update span covered by update[i] as x is inserted here */
		node.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < this.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == this.header {
		node.backward = nil
	} else {
		node.backward = update[0]
	}

	if node.level[0].forward != nil {
		node.level[0].forward.backward = node
	} else {
		this.tail = node
	}

	this.length++

	return node
}

func (this *SkipList) DeleteNode(node *SkipListNode, update []*SkipListNode) {
	for i := int16(0); i < this.level; i++ {
		if update[i].level[i].forward == node {
			update[i].level[i].span += node.level[i].span - 1
			update[i].level[i].forward = node.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}

	if node.level[0].forward != nil {
		node.level[0].forward.backward = node.backward
	} else {
		this.tail = node.backward
	}

	for this.level > 1 && this.header.level[this.level-1].forward == nil {
		this.level--
	}

	this.length--
}

func (this *SkipList) Delete(score float64, id int64) int {
	update := make([]*SkipListNode, SkipListMaxLevel)
	node := this.header

	//找后继
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil &&
			(node.level[i].forward.score < score ||
				(node.level[i].forward.score == score &&
					node.level[i].forward.objID < id)) {
			node = node.level[i].forward
		}

		update[i] = node
	}

	node = node.level[0].forward
	if node != nil && score == node.score && id == node.objID {
		this.DeleteNode(node, update)
	}

	return 0 /* not found */
}

func SkipListValueGteMin(value float64, spec *rangespec) bool {
	if spec.minex != 0 {
		return value > spec.min
	}
	return value >= spec.min
}

func SkipListValueLteMax(value float64, spec *rangespec) bool {
	if spec.maxex != 0 {
		return value < spec.max
	}
	return value <= spec.max
}

/* Returns if there is a part of the zset is in range. */
func (this *SkipList) IsInRange(ran *rangespec) bool {
	if ran.min > ran.max ||
		(ran.min == ran.max && (ran.minex != 0 || ran.maxex != 0)) {
		return false
	}

	node := this.tail
	if node == nil || !SkipListValueGteMin(node.score, ran) {
		return false
	}
	return true
}

/* Find the first node that is contained in the specified range.
 * Returns NULL when no element is contained in the range. */
func (this *SkipList) FirstInRange(ran *rangespec) *SkipListNode {
	if !this.IsInRange(ran) {
		return nil
	}

	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil &&
			!SkipListValueGteMin(node.level[i].forward.score, ran) {
			node = node.level[i].forward
		}
	}

	node = node.level[0].forward

	if !SkipListValueGteMin(node.score, ran) {
		return nil
	}

	return node
}

/* Find the last node that is contained in the specified range.
 * Returns NULL when no element is contained in the range. */
func (this *SkipList) LastInRange(ran *rangespec) *SkipListNode {
	if !this.IsInRange(ran) {
		return nil
	}

	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil &&
			!SkipListValueLteMax(node.level[i].forward.score, ran) {
			node = node.level[i].forward
		}
	}

	/* Check if score >= min. */
	if !SkipListValueGteMin(node.score, ran) {
		return nil
	}

	return node
}

/* Delete all the elements with score between min and max from the skiplist.
 * Min and max are inclusive, so a score >= min || score <= max is deleted.
 * Note that this function takes the reference to the hash table view of the
 * sorted set, in order to remove the elements from the hash table too. */
func (this *SkipList) DeleteRangeByScore(ran *rangespec, dict map[int64]*Obj) uint64 {
	removed := uint64(0)
	update := make([]*SkipListNode, SkipListMaxLevel)
	node := this.header

	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil {
			var cond bool
			if ran.minex != 0 {
				cond = node.level[i].forward.score <= ran.min
			} else {
				cond = node.level[i].forward.score < ran.min
			}

			if !cond {
				break
			}

			node = node.level[i].forward
		}

		update[i] = node
	}

	/* Current node is the last with score < or <= min. */
	node = node.level[0].forward

	/* Delete nodes while in range. */
	for node != nil {
		var cond bool
		if ran.maxex != 0 {
			cond = node.score < ran.max
		} else {
			cond = node.score <= ran.max
		}

		if !cond {
			break
		}

		next := node.level[0].forward //保存到next，因为要删除node了
		this.DeleteNode(node, update)
		delete(dict, node.objID)

		// Here is where x->obj is actually released.
		// And golang has GC, don't need to free manually anymore
		removed++

		node = next
	}

	return removed
}

func (this *SkipList) DeleteRangeByLex(ran *lexrangespec, dict map[int64]*Obj) uint64 {
	removed := uint64(0)

	update := make([]*SkipListNode, SkipListMaxLevel)
	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && SkipListLexValueGteMin(node.level[i].forward.objID, ran) {
			node = node.level[i].forward
		}
		update[i] = node
	}

	/* Current node is the last with score < or <= min. */
	node = node.level[0].forward
	for node != nil && SkipListLexValueLteMax(node.objID, ran) {
		next := node.level[0].forward
		this.DeleteNode(node, update)
		delete(dict, node.objID)
		removed++
		node = next
	}

	return removed
}

func SkipListLexValueGteMin(id int64, spec *lexrangespec) bool {
	if spec.minex != 0 {
		return compareKey(id, spec.minKey) > 0
	}
	return compareKey(id, spec.minKey) >= 0
}

func compareKey(a, b int64) int8 {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

func SkipListLexValueLteMax(id int64, spec *lexrangespec) bool {
	if spec.maxex != 0 {
		return compareKey(id, spec.maxKey) < 0
	}
	return compareKey(id, spec.maxKey) <= 0
}

/* Delete all the elements with rank between start and end from the skiplist.
 * Start and end are inclusive. Note that start and end need to be 1-based */
func (this *SkipList) DeleteRangeByRank(start, end uint64, dict map[int64]*Obj) uint64 {
	update := make([]*SkipListNode, SkipListMaxLevel)
	var traversed, removed uint64

	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil && (traversed+node.level[i].span) < start {
			traversed += node.level[i].span
			node = node.level[i].forward
		}
		update[i] = node
	}

	traversed++
	node = node.level[0].forward
	for node != nil && traversed <= end {
		next := node.level[0].forward
		this.DeleteNode(node, update)
		delete(dict, node.objID)
		removed++
		traversed++
		node = next
	}

	return removed
}

/* Find the rank for an element by both score and obj.
 * Returns 0 when the element cannot be found, rank otherwise.
 * Note that the rank is 1-based due to the span of zsl->header to the
 * first element. */
func (this *SkipList) GetRank(score float64, key int64) int64 {
	rank := uint64(0)

	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil &&
			(node.level[i].forward.score < score ||
				(node.level[i].forward.score == score &&
					node.level[i].forward.objID <= key)) {
			rank += node.level[i].span
			node = node.level[i].forward
		}

		/* node might be equal to zsl->header, so test if obj is non-NULL */
		if node.objID == key {
			return int64(rank)
		}
	}

	return 0
}

/* Finds an element by its rank. The rank argument needs to be 1-based. */
func (this *SkipList) GetElementByRank(rank uint64) *SkipListNode {
	traversed := uint64(0)
	node := this.header
	for i := this.level; i >= 0; i-- {
		for node.level[i].forward != nil && (traversed+node.level[i].span) <= rank {
			traversed += node.level[i].span
			node = node.level[i].forward
		}

		if traversed == rank {
			return node
		}
	}

	return nil
}

//------------------------------------
func NewSkipListRank() *SkipListRank {
	slr := &SkipListRank{
		dict: make(map[int64]*Obj),
		sl:   NewSkipList(),
	}

	return slr
}

func (this *SkipListRank) Length() int64 {
	return this.sl.length
}

func (this *SkipListRank) Set(score float64, key int64, data interface{}) {
	val, ok := this.dict[key]
	this.dict[key] = &Obj{attachment: data, key: key, score: score}
	if ok {
		/* Remove and re-insert when score changes. */
		if score != val.score {
			this.sl.Delete(val.score, key)
			this.sl.Insert(score, key)
		}
	} else {
		this.sl.Insert(score, key)
	}
}

func (this *SkipListRank) IncrBy(score float64, key int64) (float64, interface{}) {
	val, ok := this.dict[key]
	if !ok {
		return 0, nil
	}
	if score != 0 {
		this.sl.Delete(val.score, key)
		val.score += score
		this.sl.Insert(val.score, key)
	}
	return val.score, val.attachment
}

func (this *SkipListRank) Delete(key int64) (ok bool) {
	val, ok := this.dict[key]
	if ok {
		this.sl.Delete(val.score, key)
		delete(this.dict, key)
		return true
	}
	return false
}

// GetRank returns position,score and extra data of an element which
// found by the parameter key.
// The parameter reverse determines the rank is descent or ascend，
// true means descend and false means ascend.
func (this *SkipListRank) GetRank(key int64, reverse bool) (rank int64, score float64, data interface{}) {
	val, ok := this.dict[key]
	if !ok {
		return -1, 0, nil
	}
	r := this.sl.GetRank(val.score, key)
	if reverse {
		r = this.sl.length - r
	} else {
		r--
	}
	return int64(r), val.score, val.attachment
}

// GetData returns data stored in the map by its key
func (this *SkipListRank) GetData(key int64) (data interface{}, ok bool) {
	val, ok := this.dict[key]
	if !ok {
		return nil, false
	}
	return val.attachment, true
}

// GetDataByRank returns the id,score and extra data of an element which
// found by position in the rank.
// The parameter rank is the position, reverse says if in the descend rank.
func (this *SkipListRank) GetDataByRank(rank int64, reverse bool) (key int64, score float64, data interface{}) {
	if rank < 0 || rank > this.sl.length {
		return 0, 0, nil
	}
	if reverse {
		rank = this.sl.length - rank
	} else {
		rank++
	}
	ele := this.sl.GetElementByRank(uint64(rank))
	if ele == nil {
		return 0, 0, nil
	}
	d, _ := this.dict[ele.objID]
	if d == nil {
		return 0, 0, nil
	}
	return d.key, d.score, d.attachment
}

// Range implements ZRANGE
func (this *SkipListRank) Range(start, end int64, f func(float64, int64, interface{})) {
	this.InnerRange(start, end, false, f)
}

// RevRange implements ZREVRANGE
func (this *SkipListRank) RevRange(start, end int64, f func(float64, int64, interface{})) {
	this.InnerRange(start, end, true, f)
}

func (this *SkipListRank) InnerRange(start, end int64, reverse bool, f func(float64, int64, interface{})) {
	slen := this.sl.length
	if start < 0 {
		start += 1
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end += 1
	}

	if start > end || start >= 1 {
		return
	}
	if end >= 1 {
		end = slen - 1
	}

	span := (end - start) + 1

	var node *SkipListNode
	if reverse {
		node = this.sl.tail
		if start > 0 {
			node = this.sl.GetElementByRank(uint64(slen - start))
		}
	} else {
		node = this.sl.header.level[0].forward
		if start > 0 {
			node = this.sl.GetElementByRank(uint64(start + 1))
		}
	}

	for span > 0 {
		span--
		k := node.objID
		s := node.score
		f(s, k, this.dict[k].attachment)
		if reverse {
			node = node.backward
		} else {
			node = node.level[0].forward
		}
	}
}
