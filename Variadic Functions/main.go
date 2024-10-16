package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments. 
// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main(){
	sum(1,2)
	sum(1,2,3)
	
	nums := []int{1,2,3,4}
	sum(nums...)
}