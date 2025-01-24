package main

import "fmt"

func main() {
	sliceInt := []int{1, 5, 7, 3, 9}
	sliceString := []string{"d", "c", "b", "a"}
	fmt.Println(reverse(sliceInt))
	fmt.Println(reverse(sliceString))
}

type constrantReverse interface {
	int | string
}

func reverse[T constrantReverse](slice []T) []T {
	newSlice := make([]T, len(slice))

	sliceLen := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newSlice[sliceLen] = slice[i]
		sliceLen--
	}
	return newSlice
}
