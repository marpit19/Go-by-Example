// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it

package main

import "fmt"

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := initSeq()
	
	// We call intSeq, assigning the result (a function) to nextInt. 
	// This function value captures its own i value, 
	// which will be updated each time we call nextInt.
	
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInts := initSeq()
	fmt.Println(newInts())
}