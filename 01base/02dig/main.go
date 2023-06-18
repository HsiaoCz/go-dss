package main

import "fmt"

// 使用尾递归来实现斐波那契数列
// 所谓的斐波那契数列，指后一个数是前两个数的和

func Feb(n int, a1, a2 int) int {
	if n == 0 {
		return a1
	}

	return Feb(n-1, a2, a1+a2)
}

func main() {
	fmt.Println(Feb(1, 1, 1))
	fmt.Println(Feb(2, 1, 1))
	fmt.Println(Feb(3, 1, 1))
	fmt.Println(Feb(4, 1, 1))
	fmt.Println(Feb(5, 1, 1))
}
