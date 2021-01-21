package sl

import (
	"testing"
)

var slr *SkipListRank

func init() {
	slr = NewSkipListRank()
}

func TestSkipList_GetRank(t *testing.T) {
	if slr == nil {
		t.Failed()
	}
	slr.Set(100, 2015, "test1")
	slr.Set(99, 2016, "test2")
	slr.Set(98, 2017, "test3")
	slr.Set(97, 2018, "test4")
	slr.Set(96, 2019, "test5")
	slr.Set(95, 2020, "test6")

	rank, score, extra := slr.GetRank(2017, false)
	if rank == 3 {
		t.Log("Key:", 2017, "Rank:", rank, "Score:", score, "Extra:", extra)
	} else {
		t.Error("Key:", 2017, "Rank:", rank, "Score:", score, "Extra:", extra)
	}

	rank, score, extra = slr.GetRank(-1, false)
	if rank == -1 {
		t.Log("Key:", -1, "Rank:", rank, "Score:", score, "Extra:", extra)
	} else {
		t.Error("Key:", -1, "Rank:", rank, "Score:", score, "Extra:", extra)
	}

	id, score, extra := slr.GetDataByRank(0, true)
	t.Log("GetDataByRank[RERVERSE] Rank:", 0, "ID:", id, "SCORE:", score, "Extra:", extra)
	id, score, extra = slr.GetDataByRank(0, false)
	t.Log("GetDataByRank[UNRERVERSE] Rank:", 0, "ID:", id, "SCORE:", score, "Extra:", extra)
}

func TestSkipListRank_IncrBy(t *testing.T) {
	slr := NewSkipListRank()
	for i := 2000; i < 2022; i++ {
		slr.Set(float64(i), int64(i), "Hello World")
	}

	rank, score, _ := slr.GetRank(2016, false)
	curScore, _ := slr.IncrBy(1.5, 2016)
	if score+1.5 != curScore {
		t.Error(score, curScore)
	}

	newrank, newscore, _ := slr.GetRank(2016, false)
	if newscore != curScore {
		t.Fail()
	}
	if newrank != rank+1 {
		t.Error(newrank, rank)
	}
}
