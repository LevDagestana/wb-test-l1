package main

import "fmt"

func quickSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	firstElem := nums[0]

	lessThanFirst := []int{}
	greaterThanFirst := []int{}
	equalToFirstElem := []int{}

	for _, n := range nums {

		if firstElem > n {
			lessThanFirst = append(lessThanFirst, n)
		} else if firstElem < n {
			greaterThanFirst = append(greaterThanFirst, n)
		} else {
			equalToFirstElem = append(equalToFirstElem, n)
		}

	}

	sorted := []int{}
	if len(lessThanFirst) > 0 {
		sorted = append(sorted, quickSort(lessThanFirst)...)
	}

	sorted = append(sorted, equalToFirstElem...)

	if len(greaterThanFirst) > 0 {
		sorted = append(sorted, quickSort(greaterThanFirst)...)
	}

	return sorted
}

func main() {
	nums := []int{1, 5, 23, 6, 2, 7, 14, 3, 9, 22, 11, 19, 13, 15, 24, 16, 18, 17, 4, 8, 10, 12, 20, 25}

	sorted := quickSort(nums)

	fmt.Println(sorted)
}
