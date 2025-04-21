/*
- The value cannot be changed is know as a constant.
- The "const" keyword declares the variable as "constant", which means that it is unchangeable and read-only.
Syntax:
const CONSTNAME type = value
Example:
const pi float32 = 3.14
- Constant Types
There are two types of constants:

1. Typed constants (Example: const pi float32 = 3.14)
2. Untyped constants (Example: const pi = 3.14)
*/
package main

import "fmt"
func main()  {
	const pi float32 = 3.14
	fmt.Println(pi)

}
