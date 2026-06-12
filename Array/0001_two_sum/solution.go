package main

import "fmt"

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	// 使用 make 初始化一个 map，key 是数组里的数，value 是对应的下标
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]             // 取出当前的数值
		complement := target - num // 计算我们需要的那个配对值

		// 检查这个配对值是否在 m 里面
		if oldIndex, found := m[complement]; found {
			// 如果在记事本里找到了配对值！
			// 就去m里把那个配对值的“旧下标”查出来
			// 成功找到答案，组合【旧下标, 当前下标】，立刻返回
			return []int{oldIndex, i}
		}
		m[num] = i

	}

	return nil
}

func main() {
	// 自动推导 nums 为 []int 类型，target 为 int 类型
	nums := []int{2, 7, 11, 15}
	target := 26

	// 调用你写好的算法
	result := twoSum(nums, target)
	if result == nil {
		fmt.Println("很遗憾，数组中不存在满足条件的两个数！")
	} else {
		// 只有在 result 不为 nil 的安全情况下，才能通过索引 [0] 和 [1] 取值
		fmt.Printf("找到了！下标分别为: [%d, %d]\n", result[0], result[1])
	}
}
