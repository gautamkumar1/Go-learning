/*
A struct is a custom type in Go that lets you group related data together.
Think of it like a box that can hold multiple values with names.
Its like a blueprint for creating objects with specific properties.

💡 Why use struct?
It helps you organize data.

It's easier to work with real-world things (like users, products, cars).

You can create many objects from one struct.

Syntax:
// type StructName struct {
//     fieldName1 fieldType1
//     fieldName2 fieldType2
//     // ... more fields
// }

- Now we can also create method for struct like getData() and setData()
Syntax:
func (s StructName) getData() {
     // method code here
}
- Constructors are not available in Go, but we can create a function that initializes the struct and returns it.

*
*------------- 🧠 What is Struct Embedding? ------------
- Struct embedding means putting one struct inside another struct.
- It allows you to create complex data types by combining simpler ones.
- Organize your code better and avoid duplication.
- Sometimes works like inheritance in other languages
*

*/
package main

import (
	"fmt"
	"time"
)
type Employee struct{
	id int
	position string
	salary int
}
type Person struct {
	Name string
	Age  int
	City string
	CreatedAt time.Time // This field will be automatically set to the current time when a new Person is created

	// Another struct inside this struct
	Employee 
}
// setData
func (p *Person) setData(name string, age int, city string) {
	p.Name = name
	p.Age = age
	p.City = city
}
func (p Person) getData(Persondata Person) {
	fmt.Println(Persondata)
}
// Constructor function
func NewPerson(name string, age int, city string) *Person {
	// Initial steps here
	NewPerson:= Person{
		Name:     name,
		Age:      age,
		City:     city,
		CreatedAt: time.Now(), // Set the CreatedAt field to the current time
	}
	return & NewPerson
}
func main() {
	person1 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	//  You can also update values:
	person1.CreatedAt = time.Now()
	fmt.Println("Person1 Data: ", person1)
	// Acessing single field:
	fmt.Println("Person1 Name: ", person1.Name)

	// Methods Using
	person2 := Person{}
	person2.setData("Bob", 25, "Los Angeles")
	// fmt.Println("Person2 Data: ", person2)
	person2.getData(person2)

	// Constructor Using
	person3 := NewPerson("Charlie", 28, "Chicago")
	person3.getData(*person3)

	// If you want to use struct one time then create struct like this
	/*
	Snytax:
	varibaleName := struct {
		fieldName1 fieldType1} {fieldValue1}
	*/
	// Example:
	language := struct {
		Name string
		isGood bool
	} {
		Name: "Go",
		isGood: true,
		}
	fmt.Println("Language: ", language)


	/*
	* 
	* ------------- Struct Embedding -------------
	*

	*/

	person4 := Person{
		Name: "Gautam",
		Age:  35,
		City: "Delhi",
		Employee: Employee{ // We can also create seprate struct for Employee and use it here
			id: 1,
			position: "Software Engineer",
			salary: 50000,
		},
	}
	fmt.Println("Person4 Data: ", person4)
}