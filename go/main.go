package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("golang leetcode")
	removeResult := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(removeResult)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// slice slicing at the m becaus we all we want is nums1 without the extra zeros then append and sort nums2
	sort.Ints(append(nums1[:m], nums2...))
}
func removeElement(nums []int, val int) int {
	counter := 0
	for _, value := range nums {
		if value != val {
			nums[counter] = value
			counter++
		}
	}
	return counter
}
