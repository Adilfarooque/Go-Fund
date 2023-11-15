package addvalueintospecificpostion

import (
	"fmt"
)

func Position() {
	// Existing array
	existingArray := []int{1, 2, 3, 4, 5}

	insertIndex := 2
	newValue := 10

	// Shift elements to the right to make space for the new value
	existingArray = append(existingArray[:insertIndex+1], existingArray[insertIndex:]...)

	existingArray[insertIndex] = newValue

	fmt.Println("Array after insertion:", existingArray)
}
