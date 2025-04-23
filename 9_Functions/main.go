package main

import "fmt"

// Return value with single value
func add(a int, b int) int {
	return a + b
}
// return value with multiple values
func getLanguage() (string,string,bool){
	return "Go","Python",true
}
// Note - In Go Lang we always retrun two value and one return value is actual value of logic and 2nd one is error.
func main() {
	ans:=add(1,2)
	fmt.Println(ans)

	lang1,lang2,flag:=getLanguage()
	fmt.Println(lang1,lang2,flag)
}

