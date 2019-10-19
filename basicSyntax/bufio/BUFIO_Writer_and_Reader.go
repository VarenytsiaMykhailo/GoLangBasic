package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	rows := []string{
		"Hello Go!",
		"Welcome to Golang",
	}
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)  // Для создания потока вывода через буфер применяется функция-конструктор bufio.NewWriter()

	for _, row := range rows {
		writer.WriteString(row)  // запись строки. Принимает строку и записывает ее в виде последовательности байтов в кодировке UTF-8
		writer.WriteString(`\n`) // явный перевод строки в тексте .txt файла
		writer.WriteString("\n") // будет невидим в .txt файле
	}
	writer.Flush()// При выполнении различных методов объекта writer данные вначале накапливаются в буфере, а чтобы сбросить их в источник данных, необходимо вызвать метод Flush().

	file, err = os.Open("text.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)  // Для создания потока ввода через буфер применяется функция-конструктор bufio.NewReader()
	for {//reader.ReadString будет читать текст с того места, на котором остановился
		line, err := reader.ReadString('\n') //читает (или строго говоря, декодирует) последовательность двоичных байтов, как текст в кодировке UNF-8 или в ASCII , до байта с указанным значением, включая его ('\n') или до конца файла
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}
