/*

*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Drink struct {
	Name string
	Ice  bool
}

type Movie struct {
	Name string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
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
	
	fmt.Println("================================================")
	mm := Movie{
		Name: "Lost in Space" ,
		Rating: 9.5           ,  
	}

	fmt.Println(mm.summary())
}
