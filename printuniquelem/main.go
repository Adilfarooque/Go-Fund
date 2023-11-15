package main

import "fmt"

func main() {
	array := []int{1, 2, 2, 3, 1,9,10,11}

	for i := 0; i < len(array); i++ {
		count := 0
		for j := 0; j < len(array); j++ {
			if i!=j && array[i]==array[j] {
				count++
			}
		}
		if count == 0 {
			fmt.Println(array[i])
		}
	}
}
