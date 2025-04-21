/*
- Arrays is a collection of elements of the same type.
- Fixed Size so it is predictable.
- Memory Optimization
- Constant time access to elements. means that you can access any element in the array in constant time.
- By default Values of array
- if array is int type then default value is 0.
- if array is string type then default value is ""(empty string).
- if array is bool type then default value is false.
- if array is float type then default value is 0.0.
- if array is complex type then default value is 0+0i.
->> 2D Arrays
*/
package main

import "fmt"

func main() {
	var arr [4] int;
	// Len is used to get the length of the array.
	fmt.Println(len(arr))
	// Assigning values to the array
	arr[0] = 1;
	fmt.Println(arr)
	// 2D Array
	 arr2d:= [2][3]int{{1,2,3},{4,5,6}};
	 fmt.Println(arr2d)
	 
}