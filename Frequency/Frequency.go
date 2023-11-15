package main

import "fmt"

//Write a program in C to count the frequency of each element of an array.
//Test Data :


func countFrequency(array []int) map[int]int {
	//array := []int{1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	frequency := map[int]int{}
	for _, element := range array {
	  frequency[element] = frequency[element] + 1
	}
	return frequency
  }
  

func main(){
	array := []int{1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	frequency := countFrequency(array)
	fmt.Println(frequency)
}