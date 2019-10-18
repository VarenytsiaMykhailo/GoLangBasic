package main

import "fmt"

var stuff = "not ready"

func init() {
	stuff = "ready"
}

func main()  {
	fmt.Println("The stiff is", stuff)
}