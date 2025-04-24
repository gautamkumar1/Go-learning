/*
🧠 What is an Enum?
An enum is a set of named constants.

Example:

Small, Medium, Large (for T-shirt size)

Pending, Approved, Rejected (for status)

Go doesn't have built-in enums like other languages,
but we simulate enums using const and iota.

🚀 What is iota?
iota is a keyword that lets Go automatically count numbers starting from 0.
*/
package main

import "fmt"

type OrderStatus int

const (
	Pending OrderStatus = iota // 0
	Approved                   // 1
	Rejected                   // 2
	Cancelled                  // 3
)

func changedOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changedOrderStatus(Pending)
}