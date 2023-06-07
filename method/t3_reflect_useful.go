/*
It is not possible to compare structs of different types, and attempting to do so will result in a
compile time error. As such, it can be necessary to check that structs are of the same type before
attempting a comparison. Go’s reflect package supports this and allows the underlying type of
a struct to be interrogated. Listing 7.8 shows the reflect package being used to show the
underlying type of a struct being shown. 

Difference between keywords "new" and "make" in GOLANG: 
The new keyword in Golang is used to create a new instance of a variable, 
and it returns a POINTER to the memory allocated. 
The make keyword in Golang is used to create slices, maps, and channels, 
and it returns a value (variable?) of the type that was created.
*/
package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	fmt.Printf("%+v\n",reflect.TypeOf(b))
	fmt.Printf("%+v\n",reflect.TypeOf(a))
	fmt.Printf("%+v\n", a==b )
	fmt.Println("===================================")
	c := Drink{}
	a = b
	fmt.Printf("%+v  %+v  %p  %p \n",reflect.TypeOf(c),  c,  &a,  &b)
	fmt.Println("===================================")
	d := b
	e := b
	fmt.Printf("%p  %p  %p \n",  &b,  &d,  &e)   // até aqui cada qual com sua independencia de existencia
	fmt.Println("================  assigned by referencing a pointer. ===================")
	f := &a  // tipo de f é "*main.Drink" e o tipo de a é "main.Drink"
	f.Ice = false
	fmt.Printf("  a: %+v  f: %+v\n  &a: %p  &f: %p\n  a: %+v  f: %+v\n",  
	            a, f,  &a,  &f, reflect.TypeOf(a) , reflect.TypeOf(f) )   
				// clonou a no f ... e vice versa 

}
