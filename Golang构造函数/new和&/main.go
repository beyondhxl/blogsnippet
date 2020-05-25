// Case 1:
/*
package main

import (
	"fmt"
)

type Drink struct {
	Name    string
	Flavour string
}

func main() {
	a := new(Drink)
	a.Name = "Maaza"
	a.Flavour = "Mango"
	b := a
	fmt.Println(&a)
	fmt.Println(&b)
	b.Name = "Frooti"

	fmt.Println(a.Name)

} //This will output Frooti for a.Name, even though the addresses for a and b are different.
*/

//Case 2:
package main

import (
	"fmt"
)

type Drink struct {
	Name    string
	Flavour string
}

func main() {
	a := Drink{
		Name:    "Maaza",
		Flavour: "Mango",
	}

	b := &a
	fmt.Println(&a)
	fmt.Println(&b)
	b.Name = "Froti"

	fmt.Println(a.Name)

} //This will output Maaza for a.Name. To get Frooti in this case assign b:=&a.
