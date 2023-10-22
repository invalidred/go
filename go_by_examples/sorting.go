package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{3, 1, 9, 2}
	slices.Sort(nums)
	fmt.Println("nums", nums)

	strs := []string{"g", "c", "z", "a"}
	slices.Sort(strs)
	fmt.Println("strs", strs)

	s := slices.IsSorted(nums)
	fmt.Println("IsSorted", s)
}
