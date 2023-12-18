package main

import "fmt"


//Go struct embedding allows for structs to be embedded into interfaces
// to express a more seamless composition of types.
type base struct {
	num int
}

func (b base) describe() string{
	return fmt.Sprintf("base with num=%v",b.num)
}

//a container embeds a base
//an embedding looks like a field without a name
type container struct {
	base
	str string
}


func main(){

		//when creating structs with literals, we have to
		//initalize the embedding explicitly;
		// here the embedded types serve as the field name
		co := container{
		base: base {
			num: 1,
		},
		str: "some name",
	}

	//we can access the base's fields directly on co, e.g. co.num
	fmt.Printf("co={num: %v str: %v}\n", co.num, co.str)

	//alternativly we can spell out the full path using the embedded type name
	fmt.Println("also num:", co.base.num)
	//since container embeds base, the methods of base also become the methods of container
	//here we invoke a method that was embedded from base directly on co
	fmt.Println("describe:", co.describe())

	//Embedding structs with methods may be used
	// to bestow interface implementations onto other structs.
	// here we see that container now implements the describer interface because it
	// embeds base
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}