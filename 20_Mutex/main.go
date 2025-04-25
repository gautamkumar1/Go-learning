/*
🧠 What is a Mutex?
	- A Mutex (short for Mutual Exclusion) is like a lock.

	- It’s used when multiple goroutines are accessing or changing the same variable.

	- Only one goroutine at a time can enter the locked part — the others must wait.

🎯 Why use Mutex?
	- Without a lock, goroutines might mess each other up by writing/reading the same data at the same time — that causes bugs or weird behavior.
🚗 Real-Life Example:
	Imagine you and your friend are writing in the same notebook at the same time.
	You both scribble over each other. 🙈

	With a Mutex, only one of you writes at a time — no mess! 📝✅

*/
package main

import (
    "fmt"
    "sync"
)

var (
    counter = 0
    mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer func ()  {
		wg.Done()
		mutex.Unlock()   // 🔓 Unlock after done
	}()
    mutex.Lock()           // 🔒 Lock before accessing shared data
    counter = counter + 1
           
    
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}


