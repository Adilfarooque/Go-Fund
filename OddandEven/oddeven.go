package oddandeven

import "fmt"

//Write a program in C to separate odd and even integers into separate arrays.
func OddOrEven() {
	var Odd, Even [15]int
	var k, j int
	arr := [15]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			Even[j] = arr[i]
			j++
		} else {
			Odd[k] = arr[i]
			k++
		}
	}

	for i := 0; i < k; i++ {
		fmt.Println("", Odd[i])
	}
	for i := 0; i < j; i++ {
		fmt.Println("", Even[i])
	}
}
