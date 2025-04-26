/*
- This program demonstrates how to read and write files in Go.
- It includes functions to create a file, write to it, read from it, and delete it.
- The program also handles errors that may occur during file operations.
- 🧠 What can we do with files in Go?
	✅ Create a file

	✍️ Write to a file

	📖 Read from a file

	❌ Delete a file

We’ll use the os and io/ioutil or bufio packages for this.
🔐 Summary Table

	Action	Function/Method
	Create	os.Create("file.txt")
	Write	file.WriteString("text")
	Read	ioutil.ReadFile("file.txt")
	Append	os.OpenFile(..., os.O_APPEND, ...)
	Delete	os.Remove("file.txt")
*/
package main

func main() {
	
}