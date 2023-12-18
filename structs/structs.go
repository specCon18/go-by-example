package main

import "fmt"
// Structs are typed collections of fields.
// they're useful for grouping data together to form records
// the person struct type has name and age fields
type person struct {
	name string
	age int
}

// newPerson constructs a new person struct with the given name
// you can safely return a pointer to local variable as a local variable
// will survive the scope of the func
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main(){
	// This syntax creates a new struct
	fmt.Println(person{"Bob", 20})
	//you can name the fields when initalizing the struct
	fmt.Println(person{name: "Alice", age: 30})
	//omitted fields will be zero-valued
	fmt.Println(person{name:"Fred"})
	//an & prefix yields a pointer to the struct
	fmt.Println(&person{name:"Ann",age: 40})
	//its idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	//access struct fields with . notation
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can also use . notation with struct pointers - the pointers are
	// automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// if a struct type is only used for a single value
	// we dont have to give it a name
	// the value can have an anonymous struct type.
	// this tech is used commonly in table driven tests
	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}