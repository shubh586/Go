package main

import "fmt"
// thanks gpt
func main() {
    fmt.Println("========== Example 1: Shared Backing Array (Append Within Capacity) ==========")
    sharedBackingExample()

    fmt.Println("\n========== Example 2: New Backing Array (Append Beyond Capacity) ==========")
    newBackingExample()

    fmt.Println("\n========== Example 3: Array and Slice Sharing ==========")
    arraySliceExample()

    fmt.Println("\n========== Example 4: Array and Slice  Non backing==========")
    arraySliceNoBacking()
}

func sharedBackingExample() {
    slice1 := []int{1, 2, 3}
    slice2 := slice1[:2] // [1, 2], shares backing array

    fmt.Println("Before append:")
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)

    slice2 = append(slice2, 100) // append within original capacity
    fmt.Println("\nAfter append within capacity:")
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)

    slice2[0] = 999
    fmt.Println("\nAfter modifying slice2:")
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)
}

func newBackingExample() {
    slice1 := []int{1, 2, 3}
    slice2 := slice1[:2] // [1, 2]

    fmt.Println("Before append:")
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)

    slice2 = append(slice2, 100, 200, 300) // exceeds capacity
    fmt.Println("\nAfter append beyond capacity:")
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)

    slice2[0] = 999
    fmt.Println("\nAfter modifying slice2:")
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)
}

func arraySliceExample() {
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[:3] // [1, 2, 3]

    fmt.Println("Before modification:")
    fmt.Println("array:", arr)
    fmt.Println("slice:", slice)

    slice[1] = 100 // changes array too
    fmt.Println("\nAfter modifying slice:")
    fmt.Println("array:", arr)
    fmt.Println("slice:", slice)
}

func arraySliceNoBacking() {
    arr := [5]int{1, 2, 3, 4, 5}
    
    // Create a new slice and copy the first 3 elements
    slice :=arr[:3] // [1, 2, 3]
    

    fmt.Println("Before modification:")
    fmt.Println("array:", arr)
    fmt.Println("slice:", slice)

    slice[1] = 100 // only changes the slice, not the array

    fmt.Println("\nAfter modifying slice:")
    fmt.Println("array:", arr)  // [1 2 3 4 5] remains unchanged
    fmt.Println("slice:", slice) // [1 100 3]


	slice = append(slice, 200,300,400,600) // this will create a new backing array
	slice[0] = 999 // this will not change the array
	fmt.Println("\nAfter appending to slice:")
	fmt.Println("array:", arr)  // [1 2 3 4 5] remains unchanged
	fmt.Println("slice:", slice) // [999 100 3 200 300 400 600]
}