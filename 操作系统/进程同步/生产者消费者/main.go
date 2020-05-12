package 生产者消费者

import (
	"fmt"
	"runtime"
	"time"
)

type semaphore chan int

const maxsize = 3

var (
	empty = make(semaphore, maxsize) // 空闲缓冲区通道
	full  = make(semaphore, maxsize) // 满缓冲区通道
	mutex = make(semaphore, 1)       // 控制缓冲区的互斥访问
	items = make(semaphore, maxsize) // 生产者生产的产品
)

// 原子操作
// P操作
func (sema *semaphore) P() {
	for {
		if len(*sema) > 0 {
			<-(*sema)
			break
		}
		runtime.Gosched()
	}
}

// V操作
func (sema *semaphore) V() {
	for {
		if len(*sema) < cap(*sema) {
			*sema <- 1
			break
		}
		runtime.Gosched()
	}
}

// 生产者
func producer() {
	for {
		empty.P()
		mutex.P()
		// 生产产品
		items.V()
		fmt.Printf("[生产者] Produce One Item. Cur Items %v \n", len(items))
		mutex.V()
		full.V()
	}
}

// 消费者
func consumer() {
	for {
		full.P()
		mutex.P()
		items.P()
		fmt.Printf("[消费者] Consume One Item. Cur Items %v \n", len(items))
		mutex.V()
		empty.V()
	}
}

// 默认初始化函数
func init() {
	// 开始缓冲区是没有数据的
	for i := 0; i < maxsize; i++ {
		empty <- 1
	}
	mutex <- 1
}

func main() {
	// 生产者生产
	go producer()

	// 消费者消费
	go consumer()

	time.Sleep(time.Duration(3) * time.Millisecond)
	return
}
