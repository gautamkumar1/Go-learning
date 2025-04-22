/*
- Slice is a similar to array but it is more flexible and powerful. because we can change the size of slice at runtime.
- Slices are used to store collections of elements of the same type.
- In one word we can say that is dyanmic array.

-In Go, there are several ways to create a slice:

1. Using the []datatype{values} format
2. Create a slice from an array
3. Using the make() function

1. Create a Slice With []datatype{values}
- This is the most common way to create a slice.
Syntax:
slice_name := []datatype{values}

2. Create a Slice From an Array
- This is another way to create a slice.
Syntax:
slice_name := array_name[start_index:end_index] 

3. Using the make() function
- This is another way to create a slice.
Syntax:
slice_name := make([]datatype, length, capacity)
Note: If the capacity parameter is not defined, it will be equal to length.

-----> Memory Efficiency <-----
- When using slices, Go loads all the underlying elements into the memory.

If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 

Syntax:
copy(destination_slice, source_slice)
- The copy() function returns the number of elements copied.
*/
package main

import "fmt"

func main()  {
	// 1. Create a Slice With []datatype{values}
	mySlice := []int{}; // Create an empty slice of integers - The code above declares an empty slice of 0 length and 0 capacity.

	fmt.Println(mySlice)
	// To add elements to a slice, we can use the append() function.
	mySlice = append(mySlice, 1, 2, 3, 4, 5) // Add elements to the slice
	fmt.Println(mySlice) // Output: [1 2 3 4 5]
	// To initialize the slice during declaration, use this:
	mySlice2 := []int{1, 2, 3, 4, 5,6} // Create a slice with initial values
	fmt.Println(mySlice2) // Output: [1 2 3 4 5]
	length:= len(mySlice2) // Get the length of the slice (the number of elements in the slice)
	fmt.Println(length) // Output: 6
	capacity:= cap(mySlice2) // Get the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Println(capacity) // Output: 6

	// 2. Create a Slice From an Array
	var myArray = [5]int{1, 2, 3, 4, 5} // Create an array of integers
	arraySlice := myArray[0:5] // Create a slice from the array (from index 1 to index 4)
	fmt.Println(arraySlice) // Output: [2 3 4]

	// 3. Using the make() function
	mySlice3 := make([]int,5,10) // Create a slice of integers with length 5 and capacity 10
	fmt.Println(mySlice3) // Output: [0 0 0 0 0]
	fmt.Println("Length:", len(mySlice3)) // Output: Length: 5
	fmt.Println("Capacity:", cap(mySlice3)) // Output: Capacity: 10

	// with omitted capacity
	mySlice4 := make([]int,5) // Create a slice of integers with length 5 and capacity 5
	fmt.Println(mySlice4) // Output: [0 0 0 0 0]
	fmt.Println("Length:", len(mySlice4))
	fmt.Println("Capacity:", cap(mySlice4))
	
	// Append One Slice To Another Slice
	var slice1 = []int{1,2,3};
	var slice2 = []int{4,5,6};
	slice3 := append(slice1, slice2...) // Append slice2 to slice1
	fmt.Println(slice3) // Output: [1 2 3 4 5 6]
	// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

	// Change The Length of a Slice
	var beforeChangeLenArr = []int{1,2,3,4,5};
	fmt.Println("Before change length:", len(beforeChangeLenArr)) // Output: 5
	fmt.Println("Before change capacity:", cap(beforeChangeLenArr)) // Output: 5
	fmt.Println(beforeChangeLenArr) // Output: [1 2 3 4 5]
	// Change the length of the slice to 3
	var afterChangeLenArr = beforeChangeLenArr[:3] // Change the length of the slice to 3
	fmt.Println("After change length:", len(afterChangeLenArr)) // Output: 3
	fmt.Println("After change capacity:", cap(afterChangeLenArr)) // Output: 5
	fmt.Println(afterChangeLenArr) // Output: [1 2 3]


	// Copy - Memory Efficiency
	var numbers = []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Before copy:", numbers) // Output: [1 2 3 4 5 6 7 8 9 10]
	fmt.Println("Before copy length:", len(numbers)) // Output: 10
	fmt.Println("Before copy capacity:", cap(numbers)) // Output: 10

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-5] // Create a slice with only the first 5 elements
	numbersCopy := make([]int, len(neededNumbers)) // Create a new slice with the same length as neededNumbers
	copy(numbersCopy, neededNumbers) // Copy the elements from neededNumbers to numbersCopy
	fmt.Println("After copy:", numbersCopy) // Output: [1 2 3 4 5]




}