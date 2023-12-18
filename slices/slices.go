package main

import (
	"fmt"
	"slices"
)

func main(){
	// Unlike arrays, slices are typed only by the elements they contain.
	// not the number of elements.
	// an uninitalized slice equals to nil and has a len of 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// to create an emply slice with non-zero length,
	// use the builtin make. here we make a slice of strings of length 3 (initally zero-valued)
	// by default a new slice's capacity == its len if we know the slice is going to grow ahead of time
	// its possible to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, len(s), "cap:", cap(s))

	//setting and getting vars works just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// len reutrns the length of the slice as expected.
	fmt.Println("len:", len(s))
	// in addition to these basic opertaions,
	// slices support several more than make them richer than arrays.
	// One is the builtin append which returns a slice containing
	// one or more new values. Note that we need to accept a return value
	//from append as we make get a new slice value. 
	s = append(s, "")
	s = append(s, "e", "f")

	// slices can be copy'd
	// below we create an empty slice c
	// of the same len as s and copy into c from s
	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	//slices also support a slice operator with the syntax slice[low:high].
	// for example this gets a slice of the elements s[2],s[3],s[4]
	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	//single line declaration and initalization of a slice
	t := []string{"g","h","i"}

	//the slices package has plenty of useful funcs for slices like .Equal
	t2 := []string{"g","h","i"}
	if slices.Equal(t,t2){
		fmt.Println("t == t2")
	}

	// Slices can be composed into multi-dimensional data structues,
	// the length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i ++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}