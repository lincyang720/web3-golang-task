package main

/**
* 定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
* 然后在主函数中调用该函数并输出修改后的值。
 */
func add_pointer(a *int) int {
	if a == nil {
		return *a
	}
	return *a + 10
}

/**
* 定义一个函数，该函数接收一个整数切片作为参数，在函数内部将切片中的每个元素乘以2，
* 然后在主函数中调用该函数并输出修改后的切片。
 */
func slice_pointer(a []*int) []*int {
	if a == nil {
		return nil
	}
	for _, v := range a {
		*v = *v * 2
	}
	return a
}

func main() {
	// 	a := 10
	// 	println(add_pointer(&a))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 将 int 切片转换为 *int 切片
	pointers := make([]*int, len(a))
	for i := range a {
		pointers[i] = &a[i]
	}
	result := slice_pointer(pointers)
	for _, v := range result {
		println(*v)
	}
}
