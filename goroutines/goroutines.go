package main

import (
	"fmt"
	"time"
)

func f (from string){
	for i := 0; i<3; i++ {
		fmt.Println(from,":",i)
	}
}

func main(){
	//suppose we have function call f(s)
	//here is how wed call that in the usual way
	//running it synchronously
	f("direct")

	//To invoke this function in a goroutine, use go f(s)
	//This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	//you can also start a goroutine for an anonymous function call
	go func (msg string) {
		fmt.Println(msg)
	}("going")

	//our two function calls are running async in seperate goroutines now.
	//wait for them to finish (for a most robust approach use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}