package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main()  {
	inFileName, outFileName, err := filenamesFromCommandline()
	if err != nil {
		fmt.Println(err)
	}
	inFile, outFile := os.Stdin, os.Stdout
	if inFileName != "" {
		if inFile, err = os.Open(inFileName); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	if outFileName != "" {
		if outFile, err = os.Create(outFileName); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}
	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

func filenamesFromCommandline() (inFileName, outFileName string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		inFileName = os.Args[1]
		if len(os.Args) > 2 {
			outFileName = os.Args[2]
		}
	}
	if inFileName != "" && inFileName == outFileName {
		log.Fatal("won't overwrite the infile")
	}
	return inFileName, outFileName, nil
}

var britishAmerican = `D:\GOPROJECTS\src\github.com\VarenytsiaMykhailo\GoLangBasic\MarkSummerfieldProgrammingInGO\americanise\british-american.txt` //здесь содержатся строки с оригиинальными и замещающими их словами
func americanise(inFile io.Reader, outFile io.Writer) (err error) {
	//Для чтения из источника данных и записи в него через буфер в пакете bufio определены типы Reader и Writer.
	reader := bufio.NewReader(inFile) // Для создания потока ввода через буфер применяется функция-конструктор bufio.NewReader()
	writer := bufio.NewWriter(outFile) // Для создания потока вывода через буфер применяется функция-конструктор bufio.NewWriter()
	defer func() {
		if err == nil {
			err = writer.Flush() // При выполнении различных методов объекта writer данные вначале накапливаются в буфере, а чтобы сбросить их в источник данных, необходимо вызвать метод Flush().
		}
	}()
	var replacer func(string) string //переменная типа функция (хранит ссылку на функцию с таким прототипом)
	if replacer, err = makeReplacerFunction(britishAmerican); err != nil { // replacer - функция, отыскивающая слова, которые нужно заменить (содержатся в british-american.txt)
		return err
	}
	wordRx := regexp.MustCompile("[A-Za-z]+") //этому регулярному выражению соответствует последовательность из алфавитных символов латиницы
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n') //читает (или строго говоря, декодирует) последовательность двоичных байтов, как текст в кодировке UTF-8 или в ASCII , до байта с указанным значением, включая его ('\n') или до конца файла
		if err == io.EOF {
			err = nil //признак io.EOF не является ошибкой
			eof = true // преращаем цикл в след. итерации
		} else if err != nil {
			return err //в случае настоящей ошибки выйти немедленно
		}
		line = wordRx.ReplaceAllStringFunc(line, replacer)/* изменяем считанную строку. Метод ReplaceAllStringFunc принимает строку и функцию "реализующую замену",
		вызывает эту функцию для каждого найденного совпадения, передает этой функции совпавший текст
		и замещает совпадение текстом, возвращаемым функцией ReplaceAllStringFunc*/
		if _, err = writer.WriteString(line); err != nil { // принимает строку и записывает ее в виде последовательности байтов в кодировке UTF-8
			return err
		}
	}
	return nil
}

/*makeReplacerFunction предполагает, что файл содержи текст в кодировке UTF-8,
где в каждой строке находятся оригинальное слово и слово замены,разделенные пробельными символами*/
func makeReplacerFunction(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file) //высокоуровневая функция. Читает файл целиком и возвращает все его содержимое в виде последовательности двоичных байтов ([]byte)
	if err != nil {
		return nil, err
	}
	text := string(rawBytes)
	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line) //аналог strings.Split, разбиение идет по пробельным символам (табуляция и тд)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}
	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}






