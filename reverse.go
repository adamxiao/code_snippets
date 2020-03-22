package main

import (
//	"fmt"
)

func swap(a []int, b []int) {
	// assert len(a) == len(b)
	for i, _ := range a {
		a[i], b[i] = b[i], a[i]
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 逆反: AB -> BA
func Reverse(array []int, left int) int {

	//	fmt.Println("Reverse ", left, array[:left], array[left:])

	if left >= len(array) || left <= 0 {
		return 0
	}

	right := len(array) - left

	// case end: ArBr -> BrAr
	if right == left {
		swap(array[0:left], array[left:])
		return left
	}

	// case: ArA0Br -> BrA0Ar
	//         A0Ar ->   ArA0
	if left > right {
		swap(array[0:right], array[left:])
		return right + Reverse(array[right:], left-right)
	}

	// case: ArB0Br -> BrB0Ar
	//       BrB0   -> B0Rr
	if left < right {
		swap(array[0:left], array[right:])
		return left + Reverse(array[0:right], left)
	}

	return 0
}

// 逆反: AB -> BA
func Reverse2(array []int, left int) int {

	//	fmt.Println("Reverse2 ", left, array[:left], array[left:])

	if left >= len(array) || left <= 0 {
		return 0
	}

	// AB => A'B' => (A'B')' => BA
	reverse(array[:left])
	reverse(array[left:])
	reverse(array)

	return len(array)
}

func GenArray(size int) []int {
	array := make([]int, size)
	for i, _ := range array {
		array[i] = i + 1
	}
	return array
}

//func main() {
//	array := GenArray(31)
//	total := Reverse(array, 19)
//
//	// case left == right
//	//	array := GenArray(30)
//	//	total := Reverse(array, 15)
//
//	// case left > right
//	//	array := GenArray(31)
//	//	total := Reverse(array, 16)
//
//	//	array := GenArray(16)
//	//	total := Reverse(array, 1)
//
//	fmt.Println("array: %v", array)
//	fmt.Println("total: %v", total)
//}
