package main

import (
	"errors"
	"fmt"
)

//In Go its idiomatic to communicate errors via an explicit, seperate return value.
//this contrasts with the exceptions used in languages like java and Ruby
//and the verloaded single result/error value sometimes used in C.
//Go's approach makes it easy to see which functions return errors and to
//handle them using the same language constructs employed for any other non-error tasks


// by convention, errors are the last return value and have type error, a builtin interface
func f1(arg int) (int,error){
	if arg == 42 {
		//errors.New constructs a basic error value with the given error message
		return -1, errors.New("cant work with 42")
	}
	//a nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// its possible to use custom types as errors by implementing Error() method on them.
// Heres a variant on the example above that uses a custom error type
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int,error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg+3,nil
}

func main(){
	//The two loops below test out each of our error-returning
	//functions. Note that the use of an inline error check on the if line is a common idiom in Go code.
	for _,i := range []int{7,42} {
		if r,e :=f1(i); e != nil {
			fmt.Println("f1 failed:",e)
		} else {
			fmt.Println("f1 worked:",r)
		}
	}
	for _,i := range []int{7,42} {
		if r,e := f2(i);e != nil {
			fmt.Println("f2 failed:",e)
		} else {
			fmt.Println("f2 worked:",r)
		}
	}

	//if you want to programmatically use the data in a custom error, you'll need to get the error
	// as an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok:= e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}