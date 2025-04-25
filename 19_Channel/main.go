/*
🧃 What is a Channel?
  - A channel in Go is like a pipe.
  - It lets goroutines send and receive data to each other.

📦 Why use Channels?
  - Channels are used to communicate between goroutines.
  - They help synchronize the execution of goroutines.
  - To sync work between tasks.
  - To send data between goroutines.

There are two types of channels:
 1. Unbuffered channels: These channels block the sender until the receiver is ready to receive the data.
    Means at the time of sending data, it will wait until the receiver is ready to receive the data.
    - This is useful for synchronizing goroutines.
 2. Buffered channels: These channels allow sending data without blocking until the buffer is full.
    Means at the time of sending data, it will not wait until the receiver is ready to receive the data.
    - This is useful for decoupling goroutines.

Note: We can create a channel using the `make` function. and chan keyword.
🛑 Important: Channels are Blocking
  - If no one is receiving, send will wait.
  - If no one is sending, receive will wait.

- This is a blocking operation.
- To Get data then we used [<- channel_name]
- To send data then we used [channel_name <- data]
- To create a channel we used [make(chan data_type)]
- To close a channel we used [close(channel_name)]
*/
package main

import (
	"fmt"
	"time"
)

// This example for get data from channel  (Unbuffered channel)
func ProcessNums(nums chan int){
	fmt.Println("Processing number is ", <- nums)
}
// This example for send data to channel  (Unbuffered channel)
func Sum (nums chan int,num1 int,num2 int){
	ans := num1 + num2
	nums <- ans
}

// Goroutine synchronize Example (Unbuffered channel)
func task (done chan bool){
	defer func() { done <- true }() // Notify when the task is done , also it is run after the function exits
	fmt.Println("Task is running")
}


//  (Buffered channel) ---- Start Here
func emailSender(emailChan chan string,done chan bool) {
	defer func() { done <- true }() // Notify when the task is done
	for email := range emailChan {
		fmt.Println("Sending email to:", email)
		time.Sleep(1 * time.Second) // Simulate email sending delay
	}
	
}

func main() {
	// Multiple Channels Example - If we want to listin data on multiple channels then we can use this example.

	chan1 := make(chan string)
	chan2 := make(chan int)
	go func() {
		chan1 <- "Hello from channel 1"
	} ()

	go func ()  {
		chan2 <- 2
	} ()

	for i:=0; i<2; i++ {
		select{
			case msg := <- chan1:
				fmt.Println("Message from channel 1:", msg)
			case num := <- chan2:
				fmt.Println("Number from channel 2:", num)
		}
	}
/*
	// Buffered channel example
	// Create a buffered channel with a capacity of 100
	emailChan := make(chan string, 5)
	done := make(chan bool)
	go emailSender(emailChan, done) 
	for i := 0; i <5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com",i) // Send email to channel
	}
	fmt.Println("Sending emails done...")
	close(emailChan) // Close the channel to indicate no more emails will be sent
	<- done
*/
	// This all example for unbuffered channel
	/*
	done := make(chan bool) // Create a channel to notify when the task is done
	go task(done) // Start the task in a goroutine
	<-done // Wait for the task to finish
	*/

	/*
	// This example for send data to channel
	nums := make(chan int)
	// Sending data to channel
	go Sum(nums, 10, 20) 
	fmt.Println("Sum is ", <- nums)
	*/

	/*
	// This example for get data from channel
	numChans := make(chan int)
	// Sending data to channel
	go ProcessNums(numChans)
	numChans <- 10
	time.Sleep(2 * time.Second)
	*/

}