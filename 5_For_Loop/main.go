/*
- In Go Lang only "FOR" loop is used.
- There are no "WHILE" or "DO WHILE" loops in Go Lang.
- Syntax:

	for statement1; statement2; statement3 {
	   // code to be executed for each iteration
	}

- The Range Keyword
- The "range" keyword is used to iterate over elements in a collection (like an array, slice, or map).
- It returns two values: the index (or key) and the value of the element at that index (or key).
// - Syntax:
for index, value := range array|slice|map {
   // code to be executed for each iteration
}
Note : If you only want to access or print index value then you can use "_" (underscore) to ignore the value (actual value). vice versa.
// Example:
for index, _ := range array|slice|map {
   // code to be executed for each iteration
}
*/
package main

import (
	"fmt"
)

func main()  {
	// for i:=0; i<11; i++ {
	// 	fmt.Println(i)
	// }
	fruits := [3]string{"apple", "banana", "orange"}
	for ind,val:= range fruits {
		fmt.Println("Index:", ind, "Value:", val)
	}
}