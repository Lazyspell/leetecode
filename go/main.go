package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("golang leetcode")
	fmt.Println("_______________________________________________________________")
	fmt.Println("Merge Sorted Array")
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{4, 5, 6}
	merge2(nums1, 3, nums2, 3)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Remove Element")
	removeResult := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(removeResult)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Remove Duplicates from Sorted Array II")
	removeElementResult := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(removeElementResult)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Majority Element")
	majorityElementResult := majorityElement([]int{2, 4, 4, 1, 4, 4, 2})
	fmt.Println(majorityElementResult)
	fmt.Println("_______________________________________________________________")
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// slice slicing at the m becaus we all we want is nums1 without the extra zeros then append and sort nums2
	sort.Ints(append(nums1[:m], nums2...))
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	fmt.Println(i, j, k)
	for ; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
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

func removeDuplicates(nums []int) int {
	counterMap := make(map[int]int)
	counter := 0

	for _, value := range nums {
		counterMap[value]++
		if counterMap[value] <= 2 {
			nums[counter] = value
			counter++
		}
	}
	return counter
}

func majorityElement(nums []int) int {
	currentElement, elementCount := 0, 0
	for _, value := range nums {
		if elementCount == 0 {
			currentElement = value
		}

		if currentElement != value {
			elementCount--
		} else {
			elementCount++
		}
	}
	return currentElement
}
