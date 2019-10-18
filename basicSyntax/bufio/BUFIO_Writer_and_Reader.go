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
	writer := bufio.NewWriter(file)

	for _, row := range rows {
		writer.WriteString(row)  // запись строки
		writer.WriteString(`\n`) // явный перевод строки в тексте .txt файла
		writer.WriteString("\n") // будет невидим в .txt файле

	}
	writer.Flush() // сбрасываем данные из буфера в файл

	file, err = os.Open("text.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {//reader.ReadString будет читать текст с того места, на котором остановился
		line, err := reader.ReadString('\n')
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
