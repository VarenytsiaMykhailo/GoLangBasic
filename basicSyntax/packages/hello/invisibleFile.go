package hello

var invisibleTMP = "invisible" // если переменная/функция названа с маленькой буквы, то она будет видна только в своем пакете
func invisibleFunction() string {
	return "invisible function"
}