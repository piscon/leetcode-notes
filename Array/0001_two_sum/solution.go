package main

import "fmt"

func twoSum(nums []int, target int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				//使用占位符 fmt.Printf("[%d, %d]\n", i, j)，它能完美把整数塞进字符串里
				fmt.Printf("[%d, %d]\n", i, j)
			}
		}
	}
}

func main() {
	// 自动推导 nums 为 []int 类型，target 为 int 类型
	nums := []int{2, 7, 11, 15}
	target := 26

	// 调用你写好的算法
	twoSum(nums, target)
}
