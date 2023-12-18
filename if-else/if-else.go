package main

import "fmt"

func main(){

	//There are no ternary ifs in go

	//basic if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	//if doesnt need else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//You can use logical operators in if such as && and ||
	if  7%2 == 0 || 8%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	//a statement can precede conditionals;
	// any variables declared in the statement
	//are avalible in the current and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}