/*
- Maps are unordered collections of key-value pairs.
- The default value of a map is "nil".

** >>  Allowed Key Types
- The map key can be of any data type for which the equality operator (==) is defined. These include:

					- Booleans
					- Numbers
					- Strings
					- Arrays
					- Pointers
					- Structs
					- Interfaces (as long as the dynamic type supports equali
- Invalid key types are:

					- Slices
					- Maps
					- Functions

These types are invalid because the equality operator (==) is not defined for them.

- Go has multiple ways for creating maps.
1. Create Maps Using var and :=
Syntax:
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}
2. Create Maps Using make()
Syntax:
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)
*/
package main

import "fmt"

func main() {
	// ----------> Create Maps Using var and := <----------
	var a = map[string]string{"name":"gautam", "age": "23"}
	fmt.Println(a) // map[age:23 name:gautam]
	fmt.Println(a["name"]) // gautam
	fmt.Println(len(a)) // 2

	// ----------> Create Maps Using make() <----------
	var b  = make(map[string]string)
	fmt.Println(b) // map[] - Note: The make()function is the right way to create an empty map. 
	b["name"] = "gautam"
	b["age"] = "23"
	fmt.Println(b) // map[age:23 name:gautam]

	// ------> Empty Map <------
	// An empty map is a map that has no key-value pairs.
	// We can create an empty map using the make() function or by using the var keyword(map).
	var map1 = make(map[string]string)
	fmt.Println(map1) // map[]
	var map2 map[string]string
	fmt.Println(map2) // map[]
	fmt.Println(map1 == nil) // false - map1 is not nil, it is an empty map.
	fmt.Println(map2 == nil) // true - Note: A nil map behaves like an empty map when reading from it, but causes a runtime panic when writing to it.

	// ------> Remove Element from Map <------
// Removing elements is done using the delete() function
// Syntax: delete(map, key)
// Example:
	delete(a, "name")
	fmt.Println(a) // map[age:23] - Note: The delete() function does not return any value.

// -------> Iterate Over Maps <-------
// The range keyword is used to iterate over a map. The range keyword returns two values: the key and the value.
c := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
for key,val := range c{
	fmt.Println("Key:", key, "Value:", val) // Key: one Value: 1 Key: two Value: 2 Key: three Value: 3 Key: four Value: 4
	}
}
	
