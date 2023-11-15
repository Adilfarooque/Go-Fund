package findmaxmin

func MinAndMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0 // Return 0 for both min and max for an empty array
	}

	min := arr[0]
	max := arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}
