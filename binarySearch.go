package main

import "fmt"

func GenArray(size int) []int {
	array := make([]int, size)
	for i, _ := range array {
		array[i] = 2 * (i + 1)
	}
	return array
}

func BinarySearch(array []int, t int) int {
	l, u := 0, len(array)-1
	for l <= u {
		m := (l + u) / 2
		if array[m] == t {
			return m
		} else if array[m] > t {
			u = m - 1
		} else {
			//u = m + 1
			l = m + 1
		}
	}
	return -1
}

func main() {
	array := GenArray(100)
	var pos int

	for i := 0; i < 201; i++ {
		pos = BinarySearch(array, i+1)
		fmt.Printf("found [%v] pos %v\n", i+1, pos)
	}
}
