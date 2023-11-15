package sortdescending

import (
	"fmt"
	"sort"
)

func SortDescending() {
	arr := []int{4, 1, 6, 99, 3, 5}

	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[i] < arr[j] {
	// 			temp := arr[i]
	// 			arr[i] = arr[j]
	// 			arr[j] = temp
	// 		}
	// 	}
	// }

	// Define a custom sorting function that sorts the elements in descending order.
	// lessFunc:=func (i,j int)bool  {
	// 	return arr[i] > arr[j]
	// }
	sort.Slice(arr,func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Print("Descending Order:")
	fmt.Print(arr)

}
