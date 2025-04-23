/*
a closure is a function value that references variables from outside its body. This allows the function to access and manipulate those variables even after the outer function has returned.
*/
package main

import "fmt"

func outerFun() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	counter := outerFun()
	fmt.Println(counter())
	fmt.Println(counter())
}
