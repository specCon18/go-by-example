package main

import "fmt"

// here is a function that takes two ints and sums them
// then returns them as an int
// go requires explicit returns meaning it wont
// automatically return the value of the last expression
func plus(a int,b int) int {
	return a+b
}

//when you have multiple consecutive params with the same type they can be type declared as such:
func plusPlus(a,b,c int) int {
	return a+b+c
}

//call functions like you would expect
func main(){
	res := plus(1,2)
	fmt.Println("1+2=",res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3=",res)
}