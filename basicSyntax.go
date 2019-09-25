package main

func main() {

	//целые числа
	var i int = 10 // платформозависимый тип
	var bigInt int64 = 1<<32 - 1 // существуют int8, int16, int32, int64
	var autoInt = -10 // компилятор сам догадается о типе
	var unsignedInt uint = 12345 // платформозависимый тип
	var unsignedIntBig uint64 = 1<<64-1 // существуют uint8, uint16, uint32, uint64
	println(i, bigInt, autoInt, unsignedInt, unsignedIntBig)
	//return


	// числа с плавающей точкой
	var p float32 = 3.14 // сушествуют float32, float64
	println("float:", p)
	//return


	//булевые переменные
	var b = true
	println("bool:", b)
	//return


	//строки
	var str1 string = "Hello\n\t" // для строк - двойные кавычки. Одинарные - для символов.
	var str2 = "World"
	println(str1, str2)
	//return


	//бинарные данные
	var binary byte = '\x27'
	println("binary:", binary)
	//return


	//короткое объявление без указания var и типа:
	num := 42 //целый тип. Используемое имя переменной не должно быть объявлено ранее
	str := "Мир"
	println(num, str)


	//приведение типов
	println("float to int conversion")




}












