package main

import "fmt"

func main() {

	var a1 [3]int //инициализируется по-умолчанию нулями
	fmt.Println(a1, "длина:", len(a1)) //[0 0 0] длина: 3
	//return


	//размер можно задать константой.
	const size uint = 3
	var a2 [2*size]bool // Инициализируется по-умолчанию false
	fmt.Println(a2, "длина:", len(a2)) //[false false false false false false] длина: 6
	//return


	//[...]
	a3 := [...]int{1, 2, 3} //... - компилятор сам посчитает размер массива
	fmt.Println(a3, "длина:", len(a3)) //[1 2 3] длина: 3
	//return


	//двумерные массивы
	var aa [3][4]int //[3] [4]int
	aa[1][1] = 1
	fmt.Println(aa, "длина:", len(aa)) //[[0 0 0 0] [0 1 0 0] [0 0 0 0]] длина: 3
	//return

}
