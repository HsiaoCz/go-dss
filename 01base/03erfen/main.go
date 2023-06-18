package main

import "fmt"

// 二分查找
// 给定一个数列，并且有序，查找给定的
// 1,5,9,15,81,123,189,333

func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了
		return -1
	}
	// 从中间开始找
	mid := (1 + r) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间的数比目标大，从左边找
		return BinarySearch(array, target, 0, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(array, target, mid+1, r)
	}
}

func main() {
	array := []int{1, 5, 9, 15, 81, 123, 189, 333}
	target := 500
	result := BinarySearch(array, target, 0, len(array)-1)
	fmt.Println(result)

	target = 189
	result = BinarySearch(array, target, 0, len(array)-1)
}
