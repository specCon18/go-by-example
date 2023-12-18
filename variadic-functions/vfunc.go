package main

import "fmt"

// below is a func that takes in an arbitrary number of ints as args
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main(){
	sum(1,2)
	sum(1,2,3)
	nums := []int{1,2,3,4}
	sum(nums...)
}