package main

import (
	"sort"
)

// 仅在数组中出现了一次的数字
func onlyOnceNumber(num []int) int {
	if len(num) == 0 {
		return -1
	}

	var record = make(map[int]int)

	for _, n := range num {
		record[n]++
	}

	for k, v := range record {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 判断一个数是不是回文数字
func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Single digit numbers (0-9) are palindromes
	if x >= 0 && x < 10 {
		return true
	}

	// Store the original number for comparison later
	original := x

	// Variable to store the reversed number
	reversed := 0

	// Reverse the number digit by digit
	for x > 0 {
		// Get the last digit of x
		lastDigit := x % 10

		// Add the last digit to reversed number
		reversed = reversed*10 + lastDigit

		// Remove the last digit from x
		x = x / 10
	}

	return original == reversed
}

// 有效的括号
func isValid(s string) bool {
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

	// 有效字符串需满足：

	// 左括号必须用相同类型的右括号闭合。
	// 左括号必须以正确的顺序闭合。
	// 每个右括号都有一个对应的相同类型的左括号。
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && top != '(') ||
				(char == '}' && top != '{') ||
				(char == ']' && top != '[') {
				return false
			}
		}

	}
	return len(stack) == 0
}

//最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：

// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

// 提示：

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 如果非空，则仅由小写英文字母组成
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
		if prefix == "" {
			break
		}
	}
	return prefix
}

// 加一
func plusOne(digits []int) []int {
	/*
		给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

		将大整数加 1，并返回结果的数字数组。



		示例 1：

		输入：digits = [1,2,3]
		输出：[1,2,4]
		解释：输入数组表示数字 123。
		加 1 后得到 123 + 1 = 124。
		因此，结果应该是 [1,2,4]。
		示例 2：

		输入：digits = [4,3,2,1]
		输出：[4,3,2,2]
		解释：输入数组表示数字 4321。
		加 1 后得到 4321 + 1 = 4322。
		因此，结果应该是 [4,3,2,2]。
		示例 3：

		输入：digits = [9]
		输出：[1,0]
		解释：输入数组表示数字 9。
		加 1 得到了 9 + 1 = 10。
		因此，结果应该是 [1,0]。


		提示：

		1 <= digits.length <= 100
		0 <= digits[i] <= 9
		digits 不包含任何前导 0。
	*/
	// 从最右边的数字开始处理（最低位）
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++

		// 如果当前位小于10，说明没有进位，直接返回结果
		if digits[i] < 10 {
			return digits
		}

		// 如果当前位等于10，需要进位处理
		// 将当前位设为0，继续处理前一位
		digits[i] = 0
	}

	// 如果执行到这里，说明所有位都是9（如999），需要扩展数组
	// 创建一个长度+1的新数组，首位为1，其余为0
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result

}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	//给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

	// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：

	// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。

	// 返回 k 。
	if len(nums) == 0 {
		return 0
	}
	k := 1 // 初始化唯一元素的计数
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] { // 如果当前元素与前一个唯一元素不同
			nums[k] = nums[i] // 将当前元素放到唯一元素的后面
			k++               // 增加唯一元素的计数
		}
	}
	return k
}

// 合并区间
func merge(intervals [][]int) [][]int {
	// 	以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

	// 示例 1：

	// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
	// 输出：[[1,6],[8,10],[15,18]]
	// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
	// 示例 2：

	// 输入：intervals = [[1,4],[4,5]]
	// 输出：[[1,5]]
	// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

	// 提示：
	// 1 <= intervals.length <= 10^4
	// intervals[i].length == 2
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		println("interval: ", interval[0], interval[1])
	}
	var merged [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			merged = append(merged, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	merged = append(merged, []int{start, end})
	return merged
}

// 两数之和
func twoSum(nums []int, target int) []int {
	// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

	var m = make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	// num := []int{1, 2, 2, 4, 5, 4, 5, 7, 7, 9, 9, 1, 0}
	// println(onlyOnceNumber(num))

	// x := 121
	// x := -121
	// println(isPalindrome(x))

	// s := "({[]})"
	// s := "({[})"
	// println(isValid(s))

	// strs := []string{"flower", "flow", "flight"}
	// println(longestCommonPrefix(strs))

	// digits := []int{1, 2, 3}
	// digits := []int{9}
	// digits := []int{4, 3, 2, 1}
	// digits := []int{9, 9, 9, 9}
	// result := plusOne(digits)
	// for _, v := range result {
	// 	print(v, " ")
	// }

	// nums := []int{1, 1, 2, 3, 3}
	// result := removeDuplicates(nums)
	// println(result)

	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// intervals := [][]int{{1, 4}, {4, 5}}
	// merged := merge(intervals)
	// for _, interval := range merged {
	// 	println(interval[0], interval[1])
	// }

	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	println(result[0], result[1])

}
