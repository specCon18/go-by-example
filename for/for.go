package main

import "fmt"

func main(){
	i := 1

	// for with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	// Initial/Condition/After classic style loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// due to go not having any other loop for with no condition == while
	for {
		fmt.Println("loop")
		break
	}

	//continue passes to the next iter of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}