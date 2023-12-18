package main

import "fmt"

//the (int,int) in the func sig shows that the func returns 2 ints
func vals()(int,int){
	return 3,7
}

func main(){
	//here we use the 2 different values from the call with multiple assignment
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//if you only want a subset of the returned values use the blank identifier
	_,c := vals()
	fmt.Println(c)
}