package main

import "fmt"

func main() {
	//variadicFunctions()
	convertToArrayPointer()
}

func variadicFunctions() {
	showAllelements(1, 2, 3)
	showAllelements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}
	showAllelements(firstSlice...)

	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("Type : %T Value: %#v\n", newSlice, newSlice)
}

func convertToArrayPointer() {
	initialSlice := []int{1, 2}
	fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("Length : %d Capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	intArray := (*[2]int)(initialSlice)
	fmt.Printf("Type: %T Value: %#v\n", intArray, intArray)
	fmt.Printf("Length : %d Capacity: %d\n\n", len(intArray), cap(intArray))
}

func showAllelements(values ...int) {
	for _, val := range values {
		fmt.Println("Value", val)
	}
	fmt.Println()
}
