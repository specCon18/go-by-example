package main

import "fmt"

func main(){
	// Strings which can be concatinated by +
	fmt.Println("go"+"lang")
	// Strings can be iterpolated with a result
	// of an expression using fmt.Println and
	// seperating the expression and string by a ,
	// in this case the string is interpolated with
	// a sum of ints and sum of floats
	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0=",7.0/3.0)

	// Booleans and boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}