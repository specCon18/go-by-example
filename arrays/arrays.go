package main

import "fmt"

func main(){
	//here we create an array a that will hold exactly 5 ints.
	// arrays are typed and statically lengthed
	// by default an array is zero-valued which for ints is just 0 in all positions
	var a [5]int
	fmt.Println("emp:",a)

	//we can set a value at an index using the arrray[index] = value syntax
	//and get the value @ an index with the array[index] syntax
	//the builtin len returns the length of an array.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	//the following can be used to declare and initalize the valaues
	// of an array on a single line
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", twoD)
}