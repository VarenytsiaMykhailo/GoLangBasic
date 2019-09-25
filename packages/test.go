package main

import (
	"./hello" //"./hello" //подключаем пакет hello
	"./world1" //подключаем пакет world1
	"fmt"
)


func main()  {
	fmt.Println(hello.CallFromHello(), world1.CallFromWorld())
	//fmt.Println(hello.invisibleTMP) //эта переменная не видна вне своего пакета т.к. названа с маленькой буквы
	//fmt.Println(hello.invisibleFunction()) //эта функция не видна вне своего пакета т.к. названа с маленькой буквы
}
