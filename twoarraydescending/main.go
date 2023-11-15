package main




import(
	"fmt"
)

// Write a program in C to merge two arrays of the same size sorted in descending order.
// Test Data
func main(){
	var arr1 []int
	var arr2 []int
	var count1 int
	var count2 int
	fmt.Println("Enter The Size Of First Array:")
	fmt.Scan(&count1)
	for i := 0; i < count1; i++ {
		fmt.Printf("Element - %d :",i+1)
		var element int
		fmt.Scan(&element)
		arr1=append(arr1,element)
	}
	fmt.Println("Enter The Size Of Second Array:")
	fmt.Scan(&count2)
	for i := 0; i < count2; i++ {
		fmt.Printf("Element - %d :",i+1)
		var ele int
		fmt.Scan(&ele)
		arr2 = append(arr2,ele)
	}

	var arr3[]int = append(arr1,arr2...)


	for i := 0; i < len(arr3); i++ {
		for j := i+1; j <= len(arr3); j++ {
			if arr3[i] < arr3[j] {
				temp:=arr3[i]
				arr3[i]=arr3[j]
				arr3[j]=temp
			}
		}
	}
	fmt.Println("Descending order:")
    for i := len(arr3) - 1; i >= 0; i-- {
        fmt.Println(arr3[i])
    }
}