package main

import (
	"fmt"
	"strings"
)

type BitFlag int
const (
	Active BitFlag = 1 << iota //1 << 0 == 1 == ...00001 == 2^0
	Send // Неявно BitFlag (т.к. по умолчанию неинициализированная константа получает значение предыдущей) = 1 << 1 == 2 == ...00010 == 2^1
	Receive // Неявно BitFlag = 1 << 2 == 4 == ...000100 == 2^2
)

//String() изменяет стандартный вывод значений типа BitFlag с помощью функций пакета fmt (т.к. fmt будет выводить результат вызова функции String())
func (flag BitFlag) String() string {
	var flags []string
	if flag & Active == Active { // применяем маску Active. Если в переданном флаге 1ый бит справа == 1
		flags = append(flags, "Active")
	}
	if flag & Send == Send { // применяем маску Send. Если в переданном флаге 2ой бит справа == 1
		flags = append(flags, "Send")
	}
	if flag & Receive == Receive { // применяем маску Receive. Если в переданном флаге 2ой бит справа == 1
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {
		//приведение типа int(flag) необходимо, чтобы избежать бесконечной рекурсии
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|")) //выводит значение числовое проверяемого флага и активированные в нем флаги (1ые 3 бита)
	}
	return "0()"
}
func main()  {
	println(BitFlag(7)) //7 вывод без использования пакета fmt не пользуется определенным методом String() и выводит числовое значение
	fmt.Println(BitFlag(0)) //0()
	fmt.Println(Active) //1(Active)
	fmt.Println(Send) //2(Send)
	fmt.Println(Receive) //4(Receive)
	fmt.Println(BitFlag(3)) //3(Active|Send)
	fmt.Println(BitFlag(5)) //5(Active|Receive)
	fmt.Println(BitFlag(7)) //7(Active|Send|Receive)
	fmt.Println(BitFlag(16)) //0()  флаги Active|Send|Receive не установлены
}