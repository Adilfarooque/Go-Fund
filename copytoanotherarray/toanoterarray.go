package main

import "fmt"

func main() {
	var count int
	var arr,arr2[5]int
	fmt.Println("Enter The Size Of Array: ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Printf("Element - %d :",i+1)
		fmt.Scan(&arr[i])
	}
	for i := 0; i < count; i++ {
		arr2[i]=arr[i]
	}
	for i := 0; i < count; i++ {
		fmt.Print(arr2[i])
	}
}