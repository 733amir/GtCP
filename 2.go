// Sorting algorithms
package main

import (
	"log"
	"math/rand"
	"time"
)

func BubbleSort(array []int) []int {
	sortedArray := make([]int, len(array))
	for i := range array {
		sortedArray[i] = array[i]
	}

	for i := range sortedArray {
		for j := i + 1; j < len(sortedArray); j++ {
			if sortedArray[i] > sortedArray[j] {
				sortedArray[i], sortedArray[j] = sortedArray[j], sortedArray[i]
			}
		}
	}

	return sortedArray
}

// FIXME in place merge sort
func MergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	middle := len(array) / 2
	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])

	l, r := 0, 0
	sortedArray := make([]int, len(array))
	for i := range sortedArray {
		if l >= len(left) {
			sortedArray[i] = right[r]
			r++
		} else if r >= len(right) {
			sortedArray[i] = left[l]
			l++
		} else if left[l] <= right[r] {
			sortedArray[i] = left[l]
			l++
		} else {
			sortedArray[i] = right[r]
			r++
		}
	}

	return sortedArray
}

func main() {
	unsorted := make([]int, 10)
	rand.Seed(time.Now().Unix())
	for i := range unsorted {
		unsorted[i] = rand.Intn(10)
	}

	log.Print("Bubble Sort: ", unsorted, " => ", BubbleSort(unsorted))
	log.Print("Merge  Sort: ", unsorted, " => ", MergeSort(unsorted))
}
