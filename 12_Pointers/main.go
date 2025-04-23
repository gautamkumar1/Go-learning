/*
- Pointer is a variable that stores the memory address of another variable.
- Pointer is declared by using the * operator.
- Pointer is dereferenced by using the & operator.
- Pointer is used to pass a variable by reference to a function.
*/
package main

import "fmt"
// Call by Value 
func changeNum(num int)  {
	num = 10
	fmt.Println("Inside changeNum function:", num)
}

func changedNum2(num *int){
	*num = 10
	fmt.Println("Inside changeNum2 function:", *num)
}
func main()  {
	num := 5
	fmt.Println("Before changeNum function:", num)
	changedNum2(&num)
	fmt.Println("After changeNum function:", num)
}