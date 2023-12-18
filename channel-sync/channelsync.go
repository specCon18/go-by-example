package main

import (
	"fmt"
	"time"
)
//we can use channels to synchronize execution across goroutines.
//here is an example of using a blocking receive to wait for a goroutine
//to finish whenwaiting for multiple goroutines to finish you may prefer to use a WaitGroup


//This is the function we'll run in a goroutine.
//the done channel will be used to notify another
//goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//start a worker goroutine, giving it the channel to notify on.
	done <- true
}

func main(){
	done := make(chan bool,1)
	go worker(done)

	//block until we receive a notification from the worker on the channel
	<-done
}