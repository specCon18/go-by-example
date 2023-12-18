package main

import "fmt"

// This function returns another function,
// which we define anonymously in the body of the intSeq.
// The returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main(){
	//we call intSeq assigning the result(a function) to nextInt.
	// This function value captures its own i value, which will be updated
	// each time we call nextInt
	nextInt := intSeq()

	// see the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// to confirm that the state is unique to the instance create a new instance
	newInts := intSeq()
	fmt.Println(newInts())
}