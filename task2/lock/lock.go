package main

/**
* 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
* 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
 */

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func lock_example() {
	var mu sync.Mutex
	var count int
	var wg sync.WaitGroup
	wg.Add(10) // 启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	println("Count: ", count)
}

/**
*使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
*启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
 */
func atomic_example() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(10) // 启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("atomic Count:", count)
}

func main() {
	// lock_example()
	atomic_example()
}
