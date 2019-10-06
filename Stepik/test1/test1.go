package main

import "fmt"

func main() {

	var A [100]int
	for i := 0; i < len(A); i++ {
		A[i] = i
	}
	fmt.Println(A)
	slice := A[10:20]
	fmt.Println(slice, "len =", len(slice), "cap =", cap(slice))
	slice2 := slice[:30]
	fmt.Println(slice2, "len =", len(slice2), "cap =", cap(slice2))

}
