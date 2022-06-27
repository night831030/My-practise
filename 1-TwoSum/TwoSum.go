package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{9, 5, 3, 7, 4}
	target := 12
	var output []int
	output = twoSum(nums, target)
	TestAns := []int{1, 3}
	if !reflect.DeepEqual(TestAns, output) {
		fmt.Printf("func is fail, expect %d but get %d", TestAns, output)
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
