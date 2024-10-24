package main

import (
	"fmt"
	"sync"
)

func f(from string) {
	for i:=0;i<3;i++{
		fmt.Println(from, ":", i)
	}
}

func main() {
	//f("direct")


	//go f("goroutine")

	//go func(msg string){
	//	fmt.Println(msg)
	//}("going")

	//time.Sleep(time.Second)
	//fmt.Println("done")

	testwithWG()
}



func testwithWG() {
	f("direct")

	var wg sync.WaitGroup

	// start a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done() // to ensure wg.done() is called after finishing
		f("goroutine")
	}()

	// start with another goroutine
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		defer wg.Done()
	}("going")

	wg.Wait()
	fmt.Println("done")
}
