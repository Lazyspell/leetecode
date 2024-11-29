package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("golang leetcode")
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// slice slicing at the m becaus we all we want is nums1 without the extra zeros then append and sort nums2
	sort.Ints(append(nums1[:m], nums2...))
}
