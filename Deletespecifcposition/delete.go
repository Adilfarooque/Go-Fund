package deletespecifcposition

import "fmt"

func remove(slice []int, pos_i int) []int {
	return append(slice[:pos_i], slice[pos_i+1:]...)
}
func Delete() {
	//Write a program to delete an element at a desired position from an array.
	var pos int
	arr := []int{1, 23, 56, 77, 33, 67, 88}

	fmt.Println("Enter your position:")
	fmt.Scan(&pos)

	fmt.Println("Original Array:", arr)

	arr = remove(arr, pos)

	fmt.Println("Array after removing element at position", pos, ":", arr)
}
