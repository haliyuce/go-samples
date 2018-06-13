package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbersArray = getRandomIntArray(50)
	fmt.Println("Before sort:", numbersArray)
	mergeSort(numbersArray)
}

func bubleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		done := true
		for k := 1; k < len(array); k++ {
			if array[k-1] > array[k] {
				swapVal := array[k-1]
				array[k-1] = array[k]
				array[k] = swapVal
				done = false
			}
		}

		if done {
			fmt.Println("After sort:", array)
			break
		}
	}
}

func getRandomIntArray(size int) []int {
	var numbersArray = []int{}
	for i := 0; i < size; i++ {
		numbersArray = append(numbersArray, rand.Intn(size))
	}
	return numbersArray
}
