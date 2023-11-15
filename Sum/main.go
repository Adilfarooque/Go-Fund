package main

import "fmt"

func main() {
	var count int

	fmt.Println("Enter The Size Of Array:")
	fmt.Scan(&count)

	arr := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Printf("Element - %d :", i+1)
		fmt.Scan(&arr[i])
	}
	var sum int
	for i := 0; i < count; i++ {
		sum += arr[i]
	}

	fmt.Println("Sum Of All Array Elements:",sum)
}
