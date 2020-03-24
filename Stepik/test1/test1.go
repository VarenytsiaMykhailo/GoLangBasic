package main

import (
	"fmt"
	"os"
)

func main() {
	str := "asd"
	
	fmt.Fprintf(os.Stdout, str)
}

