package main

import "fmt"

func main()  {

	//при выводе на печать мапу, ее ключи и значения будут выводиться в случайном порядке и могут отличаться от примеров.


	//объявление
	var mm = map[string]string{} //[тип ключа] тип данных.
	//мап мапов:
	var mm2 = map[string] map[string]string{}
	fmt.Println(mm, mm2) //map[] map[]

	mm["test"] = "ok" //добавили элемент
	fmt.Println(mm) //map[test:ok]

	//объявление с помощью make:
	var mm3 = make(map[string]string)
	mm3["key1"] = "value1"
	fmt.Println(mm3) //map[key1:value1]

	var str string = "codabra"
	mm3[str] = "valueOfcodabra" // в качестве ключа можно передавать строку
	fmt.Println(mm3) //map[codabra:valueOfcodabra key1:value1]

	var boolMap = map[int]bool{9 : true, 10 : false}
	fmt.Println(boolMap[2]) //false Мап выдает значение по умолчанию, при обращении к несуществующему ключу
	//return


	//как проверить, если заданный ключ в мапе?
	lastName, ok := boolMap[2] //обращение к мапе по ключу всегда возвращает 2 значения: 1ое - значение ключа, 2ое - было ли значение по этому ключу инициализированно
	fmt.Println(lastName, ok) //false false    - 1ый false - значение по умолчанию от типа значения. 2ой false - результат отсутствия данного ключа в мапе
	_, ok2 := mm3["codabra"] //можно получать только инфу о наличии ключа в мапе
	fmt.Println(ok2) //true
	//return


	//удаление ключа из мапы
	delete(mm3, "codabra")
	_, ok3 := mm3["codabra"]
	fmt.Println(ok3) //false
	//return


	//копирование мапы. Функционала не существует. Только копированием циклом


	//мапы присваиваются по ссылке
	mm4 := mm
	mm4["newKey"] = "asd" //добавили ключ только в mm4, а он появился и в mm
	fmt.Println(mm4, mm) //map[newKey:asd test:ok] map[newKey:asd test:ok]
	//return


	//перебор в цикле
	for key, val := range boolMap {
		println(key, ":", val) // 9 : true \n 10 : false
	}
	//return


	a := 1
	switch a {
	case 1 :
		fmt.Println("a 1")
		fallthrough
	case 2 : fmt.Println("a 2")
	case 3: fmt.Println("a 3")
	default:
		fmt.Println("default")
	case 4: fmt.Println("a 4")



	}

}
