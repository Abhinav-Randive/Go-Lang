package main

import (
	"fmt"
)

func main() {
	// var arr1 = [4]int{1, 2, 3, 4}
	// arr2 := [5]int{5, 6, 7, 8, 9}

	// fmt.Println("Array 1: ", arr1)
	// fmt.Println("Array 2: ", arr2)

	// var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	// fmt.Print(cars)

	// prices := [3]int{10,20,30}

	// fmt.Println(prices[0])
	// fmt.Println(prices[2])

	// prices2 := [3]int{10,20,30}

	// prices2[2] = 50
	// fmt.Println(prices2)

	// arr1 := [5]int{} //not initialized
	// arr2 := [5]int{1,2} //partially initialized
	// arr3 := [5]int{1,2,3,4,5} //fully initialized

	// fmt.Println(arr1)
	// fmt.Println(arr2)
	// fmt.Println(arr3)

	// arr1 := [5]int{1:10,2:40}

	// fmt.Println(arr1)

	// arr2 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	// arr3 := [...]int{1,2,3,4,5,6}

	// fmt.Println(len(arr2))
	// fmt.Println(len(arr3))

	// myslice1 := []int{}
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))
	// fmt.Println(myslice1)

	// myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(myslice2))
	// fmt.Println(cap(myslice2))
	// fmt.Println(myslice2)

	// arr1 := [6]int{10, 11, 12, 13, 14,15}
	// myslice := arr1[2:4]

	// fmt.Printf("myslice = %v\n", myslice)
	// fmt.Printf("length = %d\n", len(myslice))
	// fmt.Printf("capacity = %d\n", cap(myslice))

	// myslice1 := make([]int, 5, 10)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	//   // with omitted capacity
	// myslice2 := make([]int, 5)
	// fmt.Printf("myslice2 = %v\n", myslice2)
	// fmt.Printf("length = %d\n", len(myslice2))
	// fmt.Printf("capacity = %d\n", cap(myslice2))

	// prices := []int{10,20,30}

	// fmt.Println(prices[0])
	// fmt.Println(prices[2])

	// prices := []int{10,20,30}
	// prices[2] = 50
	// fmt.Println(prices[0])
	// fmt.Println(prices[2])

	// myslice1 := []int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// myslice1 = append(myslice1, 20, 21)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	//   myslice1 := []int{1,2,3}
	//   myslice2 := []int{4,5,6}
	//   myslice3 := append(myslice1, myslice2...)

	//   fmt.Printf("myslice3=%v\n", myslice3)
	//   fmt.Printf("length=%d\n", len(myslice3))
	//   fmt.Printf("capacity=%d\n", cap(myslice3))

	//   arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	//   myslice1 := arr1[1:5] // Slice array
	//   fmt.Printf("myslice1 = %v\n", myslice1)
	//   fmt.Printf("length = %d\n", len(myslice1))
	//   fmt.Printf("capacity = %d\n", cap(myslice1))

	//   myslice1 = arr1[1:3] // Change length by re-slicing the array
	//   fmt.Printf("myslice1 = %v\n", myslice1)
	//   fmt.Printf("length = %d\n", len(myslice1))
	//   fmt.Printf("capacity = %d\n", cap(myslice1))

	//   myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	//   fmt.Printf("myslice1 = %v\n", myslice1)
	//   fmt.Printf("length = %d\n", len(myslice1))
	//   fmt.Printf("capacity = %d\n", cap(myslice1))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
