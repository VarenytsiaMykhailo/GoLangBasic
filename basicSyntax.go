package main

//объявление нескольких переменных в одном блоке. Т.к. они не в функции, то являются глобальными переменными уровня пакета
var (
	m1 int = 12
	m2 = "string"
	m3 = 23
)

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


	//комплексные числа
	var z complex64 = 2 + 3i //существуют complex64 и complex128
	println("complex number:", z)
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


	//значения по-умолчанию
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values:", defaultInt, defaultFloat, defaultString, defaultBool) //дефолтное значение для стринг - ""
	//return


	//бинарные данные
	var binary byte = '\x27'
	println("binary:", binary)
	//return


	//короткое объявление без указания var и типа:
	num := 42 //целый тип. Используемое имя переменной не должно быть объявлено ранее
	str := "Мир"
	println(num, str)
	//return


	//объявление нескольких переменных
	var v1, v2 string = "v1", "v2"
	println(v1, v2)
	println(m1, m2, m3) //смотреть объявление в одном блоке var (вначале программы, где глобальные переменные)
	//return


	//приведение типов
	var floatTMP float32 = 32.7
	println("float to int conversion:", int(floatTMP)) //32
	var unsignedTMP uint = 3
	var signedTMP int = 11
	println(signedTMP + int(unsignedTMP))
	println("int to string conversation:", string(65)) //буква "A" по номеру в таблице юникода
	//return


	//конкатенация строк и длина строки
	s1 := "Adam"
	s2 := "Babam"
	fullName := s1 + " " + s2
	println(fullName, len(fullName)) //len - считает длину переданной строки

	//ескейп сиволы в строках
	escaping := `Hello\r\n //comment 
	World` //используется при создании шаблонов и регулярных выражений
	println(escaping) //\r\n //comment, перевод строки и табуляцию перед World выведет на экран
	//return

}












