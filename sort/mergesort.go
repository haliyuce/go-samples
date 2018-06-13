package main

import (
	"fmt"
	"math"
)

func mergeSort(array []int) {
	// respond := make(chan []int, 20)
	lastvalue := process(array, 0, len(array)-1)
	fmt.Println(lastvalue)
}

func process(array []int, start int, end int) []int {
	length := float64(len(array))

	if 3 > len(array) {
		if len(array) == 1 {
			return array
		}

		if array[0] > array[1] {
			swapVal := array[0]
			array[0] = array[1]
			array[1] = swapVal
		}
		return array
	}

	middle := int(math.Floor(length / 2))

	return merge(process(array[0:middle], start, start+middle), process(array[middle:int(length)], start+middle+1, end))
}

func merge(firstArray []int, secondArray []int) []int {
	var firstArrayCounter int
	var secondArrayCounter int

	var lastSlice []int
	for {
		if firstArray[firstArrayCounter] < secondArray[secondArrayCounter] {
			lastSlice = append(lastSlice, firstArray[firstArrayCounter])
			firstArrayCounter++
		} else {
			lastSlice = append(lastSlice, secondArray[secondArrayCounter])
			secondArrayCounter++
		}

		if firstArrayCounter == len(firstArray) {
			lastSlice = append(lastSlice, secondArray[secondArrayCounter:len(secondArray)]...)
			break
		}

		if secondArrayCounter == len(secondArray) {
			lastSlice = append(lastSlice, firstArray[firstArrayCounter:len(firstArray)]...)
			break
		}
	}

	return lastSlice
}
