package hello //название пакета дожно совпадать с названием директории. При этом название файла может быть любым

func CallFromHello() string { //если метод начинается с большой буквы - он доступен в других пакетах
	return "HELLO" + " " + invisibleTMP //переменная invisibleTMP из файла invisibleFile видна в своем пакете
}