package main

import (
	deletespecifcposition "BasiWorld/Deletespecifcposition"
)

func main() {
	// Declare an array of size n.
	// var n int
	// fmt.Print("Enter the number of values: ")
	// fmt.Scanln(&n)
	// array := make([]int, n)

	// // Read n number of values from the user into the array.
	// for i := 0; i < n; i++ {
	// 	fmt.Print("Enter value: ")
	// 	fmt.Scanln(&array[i])
	// }
	// // Create a new empty array.
	// var reversedArray []int
	// // Iterate over the original array in reverse order, copying each element to the new array.
	// for i := n - 1; i >= 0; i-- {
	// 	reversedArray = append(reversedArray, array[i])
	// }
	// // Print the new array.
	// fmt.Println("The values in reverse order are:")
	// for i := 0; i < n; i++ {
	// 	fmt.Println(reversedArray[i])
	// }

	// r := reversedArray[:3]
	// fmt.Println("", r)

	//Find Max and Min Elements in an array

	// var count int
	// fmt.Println("Enter the Length:")
	// fmt.Scan(&count)
	// arr := make([]int, count)

	// for i := 0; i < count; i++ {
	// 	fmt.Printf("Element - %d :", i+1)
	// 	var element int
	// 	fmt.Scan(&element)
	// 	arr = append(arr, element)
	// }

	// //array := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	// if count > 0 {
	// 	min, max := findmaxmin.MinAndMax(arr)
	// 	fmt.Printf("Minimum: %d\n", min)
	// 	fmt.Printf("Maximum: %d\n", max)
	// }

	//oddandeven.OddOrEven()
	//sorteleascendingorder.AscedingOrder()
	//sortdescending.SortDescending()
	//addvalueintospecificpostion.Position()
	deletespecifcposition.Delete()
}
