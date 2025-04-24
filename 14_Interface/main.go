/*
- An interface is like a set of rules.
It says: "If a type wants to follow me, it must have these methods."

It helps you:

write flexible code

use different types with the same function

think like: "Anything that can do this is okay"

- An interface is a contract that a type must fulfill.

📦 Real-Life Example: Animal Sound
Let’s say we want different animals to make a sound.
Each animal makes a different sound, but we want to treat them the same way.

*/

package main

import "fmt"

// Define interface
type Animal interface {
    Speak() string
}
// Define types that implement the interface
type Dog struct{

}
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{

}

func (c Cat) Speak() string {
	return "Meow!"
}

// Function that takes the interface as a parameter
func MakeSound(a Animal) {
	// Call the Speak method of the interface
	sound := a.Speak()
	fmt.Println(sound)
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{}
	cat := Cat{}

	// Call the MakeSound function with different types
	MakeSound(dog) // Output: Woof!
	MakeSound(cat) // Output: Meow!
}