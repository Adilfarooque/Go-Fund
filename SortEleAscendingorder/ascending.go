package sorteleascendingorder

import (
	"fmt"
	"sort"
)

func AscedingOrder() {
	arr := []int{23, 56, 92, 88, 44}

	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[j] < arr[i] {
	// 			temp := arr[j]
	// 			arr[j] = arr[i]
	// 			arr[i] = temp
	// 		}
	// 	}
	// }

	// fmt.Print("Asecnding Order:")
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("%v ", arr[i])
	// }

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)

}
