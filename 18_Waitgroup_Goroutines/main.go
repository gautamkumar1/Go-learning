/*
🧠 What is a WaitGroup?
- When you run multiple goroutines, you sometimes want to wait for all of them to finish before moving on.

That’s exactly what sync.WaitGroup is for.
It helps your main function wait until goroutines are done.

🧾 Example:
- “I want to wait for all my tasks to finish before I leave.”
- “I want to wait for all my downloads to finish before I close the app.”
*/
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // tell WaitGroup: "I'm done!"
    fmt.Printf("Worker %d: starting\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Worker %d: done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // tell WaitGroup: "one more goroutine is coming"
        go worker(i, &wg)
    }

    wg.Wait() // wait until all goroutines call Done()
    fmt.Println("All workers finished!")
}
