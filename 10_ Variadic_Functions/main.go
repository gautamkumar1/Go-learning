// Variadic Functions - That function accepts N number of parameters is called variadic function. Ex- Println() function
package main

import "fmt"
func sum(num...int) int{
	sum :=0;
	for _,val := range num{
		sum+= val
	}
	return sum
}
func main()  {
	fmt.Println("Sum of 1,2,3,4,5 is ",sum(1,2,3,4,5)) // 15
}