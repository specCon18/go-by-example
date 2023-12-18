package main

import "fmt"

// go supports pointers allowing you to pass refs to values and records
// in your program


//this func takes in an int as a param
func zeroval(ival int){
	ival = 0
}

//this func takes in an pointer to an int as a parameter
func zeroptr (iptr *int) {
	*iptr = 0
}

func main() {
	//display initial value of i
	i := 1
	fmt.Println("initial:",i)

	//display the value of i after passing to zeroval
	zeroval(i)
	fmt.Println("zeroval:",i)

	//display the value of the data @ the memory address of i aka i's pointer
	zeroptr(&i)
	fmt.Println("zeroptr:",i)

	//pointers can be printed as well and will return their memory address
	fmt.Println("pointer:",&i)

	//zeroval doesnt change the val of i but zeroptr does because it edits the data at the memory address

}