package main

func main() {
	x, y, z := foo() //присвоение возвращаемых значений из функции foo
	println(x, y, z) //100 false abraCodabra

	//можно игнорировать некоторые возвращаемые значения с помощью имен "_":
	_, d, _ := foo()
	println(d) //false
}

func foo() (int, bool, string) {
	return 100, false, "abraCodabra" //возвращаем несколько типов значений из функции
}