/*
- Generics allow you to write functions and data structures that can work with any data type.

  - Think of it like:
    “I don’t care if it’s an int, string, or float — I’ll work with it!”
  - Generics are a way to write code that is flexible and reusable.
  - It helps you:
		- Avoid code duplication
		- Write cleaner code
- In the above code, we have two functions that do the same thing but for different data types.
	Problem: We have to write two functions for different data types.
	Solution: We can use generics to write a single function that works with any data type.
*/
package main

import "fmt"
// This is without generics
func PrintSlice (num []int){
	for _,val := range num{
		fmt.Println(val)
	}
}
func PrintSliceString (num []string){
	for _,val := range num{
		fmt.Println(val)
	}
}
// This is with generics
// Instead of any we also give Interface{} but it is not recommended as it is not type safe.
// Instead of any we can also define a type that we want to accept in the function.
func PrintSliceGeneric[T any] (num []T){
	for _,val := range num{
		fmt.Println(val)
	}
}

func main() {
	// This is without generics
	PrintSlice([]int{1, 2, 3, 4, 5})
	PrintSliceString([]string{"a", "b", "c", "d", "e"})
	/*
	- In the above code, we have two functions that do the same thing but for different data types.
	Problem: We have to write two functions for different data types.
	Solution: We can use generics to write a single function that works with any data type.
	*/
	// This is with generics
	PrintSliceGeneric([]int{1, 2, 3, 4, 5})
	PrintSliceGeneric([]string{"a", "b", "c", "d", "e"})
	
}