/*
Structs are commonly used in C and C++ and, as a C-family language, 
Golang also has structs.
Structs are not a way to create classes or object-oriented code. 
Structs are a way to create data structures that are 
relevant to the data you are trying to model.  
*/
package main

import (
	"fmt"
)

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	// var m Movie
	// fmt.Printf("%+v\n", m)
	// m.Name = "Metropolis"
	// m.Rating = 0.9918
	// fmt.Printf("%+v\n", m)

	/* "new" in Golang is a keyword that is used to instantiate a struct
	    (custom data type).
		The new keyword returns a pointer to the instantiated variable.
		It creates a "zeroed" storage for a new item.  Mar 9, 2023
	*/

	m := new(Movie) // cria um ponteiro ??? certo ???
	fmt.Printf("%+v\n", m)
	m.Name = "Metropolis 2"
	m.Rating = 0.59
	fmt.Printf("%+v\n", m)
	fmt.Println(m)

	// using SHORT VARIABLE ASSIGNMENT
	c := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}
	fmt.Println("\nesta é a 'c'",c)
	fmt.Printf("%+v\n",c)
}
