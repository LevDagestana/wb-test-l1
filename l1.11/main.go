package main

import "fmt"

func findIntersection(nums1 []int, nums2 []int) {
	var numsMap = make(map[int]int)
	var intersection []int

	for _, num := range nums1 {
		if numsMap[num] > 0 {
			continue
		} else {
			numsMap[num] = 1
		}
	}

	for _, num := range nums2 {
		if numsMap[num] > 0 {
			numsMap[num] += 1
		} else {
			continue
		}
	}

	for num, count := range numsMap {
		if count > 1 {
			intersection = append(intersection, num)
		}
	}
	fmt.Println(intersection)
}

func main() {
	nums1 := []int{3, 6, 1, 5, 8, 6}

	nums2 := []int{9, 1, 2, 3, 8, 9}

	findIntersection(nums1, nums2)
}
