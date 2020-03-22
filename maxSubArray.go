package main

import (
	"fmt"
	"math/rand"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(array []int) int {
	maxendinghere := 0
	maxsofar := 0
	l, u := -1, -1
	maxendinghere_l := 0
	for i, val := range array {
		maxendinghere = max(maxendinghere+val, 0)

		// save max sub array index [l,u]
		if maxendinghere == 0 {
			maxendinghere_l = i + 1
		}
		if maxendinghere > maxsofar {
			l = maxendinghere_l
			u = i
		}

		maxsofar = max(maxsofar, maxendinghere)
	}

	fmt.Printf("[%v,%v]\n", l, u)
	if l < 0 || l >= len(array) || l > u || u >= len(array) {
		panic("[l,u] wrong")
	}

	sum_lu := 0
	for i := l; i <= u; i++ {
		sum_lu += array[i]
	}
	if sum_lu != maxsofar {
		panic("maxsofar is not correct")
	}
	fmt.Printf("max sub array %v\n", array[l:u+1])

	return maxsofar
}

func maxSubArray2(array []int) int {
	maxsofar := 0
	for i := 0; i < len(array); i++ {
		sum := 0
		for j := i; j < len(array); j++ {
			sum += array[j] // sum [i,j]
			maxsofar = max(maxsofar, sum)
		}
	}
	return maxsofar
}

func GenArray(size int) []int {
	array := make([]int, size)
	for i, _ := range array {
		array[i] = rand.Intn(200) - 100
	}
	return array
}

func main() {
	array := GenArray(100)
	fmt.Printf("%v\n", array)

	maxsofar := maxSubArray(array)
	fmt.Printf("maxsofar %v\n", maxsofar)

	maxsofar2 := maxSubArray2(array)
	fmt.Printf("maxsofar2 %v\n", maxsofar2)

}
