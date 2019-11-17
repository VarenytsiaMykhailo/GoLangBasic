package main

import (
	"fmt"
	"github.com/VarenytsiaMykhailo/GoLangBasic/coursera/visibility/person"
)

func main() {
	//p := person.NewPerson(1, "rvasily", "secret")
	person1 := person.NewPerson(2, "Mike", "abra")
	fmt.Println(person1)
	fmt.Println(person.GetSecret(person1))
	person1.UpdateSecret("codabra")
	fmt.Println(person1)
	fmt.Println(person.GetSecret(person1))
	// p.secret undefined (cannot refer to unexported field or method secret)
	// fmt.Printf("main.PrintPerson: %+v\n", p.secret)

	//secret := person.GetSecret(p)
	//fmt.Println("GetSecret", secret)
}
