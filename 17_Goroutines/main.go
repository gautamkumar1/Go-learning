/*
🧠 What is a Goroutine?
- A Goroutine is a lightweight thread managed by Go.
- It helps you run code at the same time (concurrently).
- You can do multiple tasks together instead of one-by-one.
🧾 Example:
- “I want to print something while doing other work.”
- “I want to download a file while showing a progress bar.”
🧩 Real Life Example:
Imagine you're:

Cooking (1 task)

Talking on the phone (2nd task)

With goroutines, you can do both at the same time. 🍳📞

Go programs exit when main finishes, even if goroutines are still running.
So, we add time.Sleep(...) to wait.
*/
package main

import (
    "fmt"
    "time"
)

func sayHello(){
	for i := 0; i <5; i++ {
		fmt.Println("Hello Goroutine", i)
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds
	}
	
}

func main()  {
	go sayHello()  // Run in background (goroutine)
	for i := 0; i < 5; i++ {
		fmt.Println("Hello Main", i)
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds

	}
	
}