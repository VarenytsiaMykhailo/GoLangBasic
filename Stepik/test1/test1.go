package main

import "fmt"

func main() {
	var a, b, i int
	fmt.Scan(&a, &b)
	if (a / 1000 != 0){
		i = 1000
	} else if (a / 100 != 0){
		i = 100
	} else if (a / 10 != 0){
		i = 10
	}

	fmt.Println(i)
	/*for ; i >= 10; i = i / 10{
		switch ((a / i) % 10) {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
		}
	}*/
}
