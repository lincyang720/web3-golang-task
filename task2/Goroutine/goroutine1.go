package main

import "time"

/**
 * 创建两个goroutine，一个打印偶数，一个打印奇数
 */
// func main() {
// 	// Example usage of the functions
// 	go func() {
// 		for i := 1; i <= 10; i += 1 {
// 			if i%2 != 0 {
// 				println("Even number: ", i)
// 			}
// 		}

// 	}()
// 	go func() {
// 		for j := 2; j <= 10; j += 1 {
// 			if j%2 == 0 {
// 				println("Odd number: ", j)
// 			}
// 		}
// 	}()

// 	// 添加等待时间，确保goroutine有时间执行完成
// 	time.Sleep(100 * time.Millisecond)
// }

/**
 * 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
* 考察点 ：协程原理、并发任务调度。
*/
func taskScheduler(tasks []func()) {
	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			elapsed := time.Since(start)
			println("Task executed in: ", elapsed.Milliseconds(), "ms")
		}(task)
	}

	// 添加等待时间，确保所有任务有时间执行完成
	time.Sleep(100 * time.Millisecond)
}

func main() {
	tasks := []func(){
		func() {
			time.Sleep(50 * time.Millisecond)
			println("Task 1")
		},
		func() {
			time.Sleep(50 * time.Millisecond)
			println("Task 2")
		},
		func() {
			time.Sleep(30 * time.Millisecond)
			println("Task 3")
		},
	}

	taskScheduler(tasks)
}
