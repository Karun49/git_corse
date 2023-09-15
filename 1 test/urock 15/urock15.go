package main

import (
	"fmt"
)

func main() {
	slices()
}

func slices() {
	var defaltSlices []int
	fmt.Printf("Type : %T Value %#v", defaltSlices, defaltSlices)
	fmt.Printf("\nLenght: %d Capacity: %d\n", len(defaltSlices), cap(defaltSlices))

	stringSliceLiteral := []string{"First", "Second"}
	fmt.Printf("Type : %T Value %#v", stringSliceLiteral, stringSliceLiteral)
	fmt.Printf("\nLenght: %d Capacity: %d\n", len(stringSliceLiteral), cap(stringSliceLiteral))

	sliceByMake := make([]int, 0, 5)
	fmt.Printf("Type : %T Value %#v", sliceByMake, sliceByMake)
	fmt.Printf("\nLenght: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Printf("Type : %T Value %#v", sliceByMake, sliceByMake)
	fmt.Printf("\nLenght: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 6)
	fmt.Printf("Type : %T Value %#v", sliceByMake, sliceByMake)
	fmt.Printf("\nLenght: %d Capacity: %d\n", len(sliceByMake), cap(sliceByMake))

	for ind, value := range sliceByMake {
		fmt.Printf("Index : %d Value : %d\n", ind, value)
	}
}

/*
type Person struct {
	age  int
	name string
}

func main() {
	arrays()
}

func arrays() {
	var intarr [3]int
	fmt.Printf("%T Value: %#v\n", intarr, intarr)

	intarr[0] = 5
	intarr[1] = 6

	fmt.Printf("%T Value: %#v\n", intarr, intarr)

	people := [2]Person{
		{
			name: "Katy",
			age:  18,
		},
		{
			name: "Josh",
			age:  23,
		},
	}
	fmt.Printf("Type: %T Value: %#v\n", people, people)

	stringArr := [...]string{"First", "Second", "Third", "Fourth"}
	fmt.Printf("%T Value: %#v\n", stringArr, stringArr)
	fmt.Printf("Lenght: %d Capacity: %d\n", len(stringArr), cap(stringArr))

	for i := 0; i < len(stringArr); i++ {
		fmt.Printf("Index: %d Value: %s\n", i, stringArr[i])
	}

	for inx, value := range stringArr {
		fmt.Printf("Index: %d Value: %s\n", inx, value)
	}

	for _, value := range intarr {
		fmt.Printf("Value: %d\n", value)
	}

	newIntArr := changeArray(intarr)
	fmt.Printf("Type: %T Value: %#v\n", intarr, intarr)
	fmt.Printf("Type: %T Value: %#v\n", newIntArr, newIntArr)
}

func changeArray(arr [3]int) [3]int {
	arr[2] = 3
	return arr
}*/
