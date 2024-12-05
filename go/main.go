package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println("golang leetcode")
	fmt.Println("_______________________________________________________________")
	fmt.Println("Merge Sorted Array")
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{4, 5, 6}
	merge2(nums1, 3, nums2, 3)
	fmt.Println(nums1)
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
	fmt.Println("Fibonacci linear")
	fibResult := fib(6)
	fmt.Println(fibResult)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Fibonacci recursive")
	fibonacciResult := fibonacci(6)
	fmt.Println(fibonacciResult)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Rotate Array")
	rotateList := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(rotateList, 3)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Best Time to Buy and Sell Stock")
	maxProfitList := []int{7, 1, 5, 3, 6, 4}
	maxProfitResult := maxProfit(maxProfitList)
	fmt.Println("maxProfit:", maxProfitResult)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Product of Array Except Self")
	productExceptSelfList := []int{7, 1, 5, 3, 6, 4}
	productExceptSelfResult := productExceptSelf(productExceptSelfList)
	fmt.Println(productExceptSelfResult)
	fmt.Println("_______________________________________________________________")
	fmt.Println("Valid Palindrome")
	palindromeResult := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(palindromeResult)
	fmt.Println("_______________________________________________________________")
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m], nums2...))
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
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
	counter := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
func removeDuplicates2(nums []int) int {
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

func fib(num int) int {
	ans := []int{0, 1}

	for i := 2; i < num+1; i++ {
		ans = append(ans, ans[i-1]+ans[i-2])
	}
	fmt.Println(ans)
	return ans[num]
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

func rotate(nums []int, k int) {
	if len(nums) < k {
		k = k % len(nums)
	}
	right := len(nums) - k

	copy(nums, append(nums[right:], nums[:right]...))
}

func maxProfit(prices []int) int {
	maxProfit, left, right := 0, 0, 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			maxProfit = max(maxProfit, prices[right]-prices[left])
		} else {
			left = right
		}
		right++
	}

	return maxProfit
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	pre := 1

	for index, value := range nums {
		result[index] = pre
		pre *= value
	}

	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= post
		post *= nums[i]
	}

	return result
}
func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^0-9a-zA-Z]+`)
	clean := strings.ToLower(re.ReplaceAllString(s, ""))
	for i := 0; i < len(clean)/2; i++ {
		if clean[i] != clean[len(clean)-1-i] {
			return false
		}
	}
	return true
}
