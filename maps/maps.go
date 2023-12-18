package main

import (
	"fmt"
	"maps"
)

func main() {
	// to create an empty map, use the builtin make:
	// make(map[key-type]val-type)
	m := make(map[string]int)

	//set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// printing a map with e.g.
	// fmt.Println will show all of its key/value pairs
	fmt.Println("map:",m)

	// Get a value from a key with name[key]
	// if the key doesn't exist the zero value of the value type is returned
	
	v1 := m["k1"]
	fmt.Println("v1:",v1)

	v3 := m["k3"]
	fmt.Println("v3:",v3)

	//the builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	//the builtin delete removes key/value pairs from a map,
	//use the clear builtin to remove all key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:",m)

	clear(m)
	fmt.Println("map:",m)

	// the optional second return value when getting the value
	// from a map indicates if the key was present in the map allowing
	// to diferentiate between default zero-values and values acutally stored as 0
	// here we dont need the value itself so we ignore it with the _ or blank identifier
	_, prs := m["k2"]
	fmt.Println("prs:",prs)

	//you can declare and initalize a new map in the same line with the following syntax:
	n := map[string]int{"foo": 1,"bar": 2}
	fmt.Println("map:", n)

	//the map package also contains useful funcs for map such as .Equal
	n2 := map[string]int{"foo": 1,"bar":2}
	if maps.Equal(n,n2){
		fmt.Println("n == n2")
	}
}