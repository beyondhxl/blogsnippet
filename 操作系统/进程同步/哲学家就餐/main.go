package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Gosched yields the processor, allowing other goroutines to run.
It does not suspend the current goroutine, so execution resumes automatically.
*/

type semaphore chan int

// 全局变量
var (
	chopsticks = make([]semaphore, 5) // 5个哲学家
	mutex      = make(semaphore, 1)   // 互斥锁
)

func (sema *semaphore) P() {
	for {
		if len(*sema) > 0 {
			<-*sema
			break
		}
		runtime.Gosched()
	}
}

func (sema *semaphore) V() {
	for {
		if len(*sema) < cap(*sema) {
			*sema <- 1
			break
		}
		runtime.Gosched()
	}
}

// 进餐
func dining(i int) {
	for {
		// 上锁
		mutex.P()
		// 取左边的筷子
		chopsticks[i].P()
		// 取右边的筷子
		chopsticks[(i+1)%5].P()
		// 解锁
		mutex.V()
		fmt.Printf("Philosopher %v is eating \n", i+1)
		// 放下左边的筷子
		chopsticks[i].V()
		// 放下右边的筷子
		chopsticks[(i+1)%5].V()
		fmt.Printf("Philosopher %v is thinking \n", i+1)

	}
}

func init() {
	for i := 0; i < 5; i++ {
		chopsticks[i] = make(semaphore, 1)
		chopsticks[i].V()
	}
	mutex.V()
}

func main() {
	for i := 0; i < 5; i++ {
		go dining(i)
	}

	// 主线程阻塞一段时间
	time.Sleep(time.Duration(3) * time.Millisecond)
	return
}
