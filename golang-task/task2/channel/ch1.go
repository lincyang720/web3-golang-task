package main

import "time"

/**
*编写一个程序，使用通道实现两个协程之间的通信。
* 一个协程生成从1到10的整数，并将这些整数发送到通道中，
*另一个协程从通道中接收这些整数并打印出来。
 */
func channel_example() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			println("receive data from channel: ", num)
		}
	}()
}

func channel_example2() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			println("receive data from channel: ", num)
		}
	}()
}

func main() {
	// channel_example()
	channel_example2()
	// 添加等待时间，确保所有协程有时间执行完成
	time.Sleep(100 * time.Millisecond)
}
