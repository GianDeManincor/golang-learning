package main

import "fmt"

func main() {
	sliceInt := []int{1, 5, 7, 3, 9}
	fmt.Println(reverse(sliceInt))
}

func reverse(slice []int) []int {
	newSliceInt := make([]int, len(slice))

	sliceLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newSliceInt[sliceLen] = slice[i]
		sliceLen--
	}
	return newSliceInt
}
