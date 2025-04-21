/*
- Variables are containers for storing data values.
- Can be used inside and outside of functions
- There are four types of variables in Go:
1. Int
2. String
3. Float
4. Bool

-- Declaring (Creating) Variables --
1. Using var keyword with Variable name and type
Syntax: var variablename type = value
var a int = 10
Note: You always have to specify either type or value (or both).
2. With the := sign:
Syntax: variablename := value
var a = 10
Note:
- This is shorthand for declaring and initializing a variable.
- means that the compiler decides the type of the variable, based on the value
- It is not possible to declare a variable using ":=" without assigning a value to it.
- Can only be used inside functions

*/
package main

import "fmt"
func main() {
	var str string =  "gautam";
	var num int = 10;
	var isFalg bool = true;
	var pi float32 = 3.14;
	fmt.Println(str,num,isFalg,pi);
}