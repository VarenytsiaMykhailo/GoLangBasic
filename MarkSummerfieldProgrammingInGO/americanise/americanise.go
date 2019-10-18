package main
//D:\GOPROJECTS\src\github.com\VarenytsiaMykhailo\GoLangBasic\MarkSummerfieldProgrammingInGO\americanise\americanise.go
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main()  {
	inFileName, outFileName, err := filenamesFromCommandline()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("inFileName and outFileName: ", inFileName, outFileName) // отладка
	inFile, outFile := os.Stdin, os.Stdout
	fmt.Println("inFile and outFile: ", inFile, outFile) // отладка
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
	fmt.Println("_____________________________________________________") //отладка
	fmt.Println("inFileName and outFileName: ", inFileName, outFileName) // отладка
	fmt.Println("inFile and outFile: ", inFile, outFile) // отладка

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

var britishAmerican = "british-american.txt"
func americanise(inFile io.Reader, outFile io.Writer) (err error) {
	//Для чтения из источника данных и записи в него через буфер в пакете bufio определены типы Reader и Writer.
	reader := bufio.NewReader(inFile) // Для создания потока ввода через буфер применяется функция bufio.NewReader()
	writer := bufio.NewWriter(outFile) // Для создания потока вывода через буфер применяется функция bufio.NewWriter()
	defer func() {
		if err == nil {
			err = writer.Flush() // При выполнении различных методов объекта writer данные вначале накапливаются в буфере, а чтобы сбросить их в источник данных, необходимо вызвать метод Flush().
		}
	}()
	var replacer func(string) string //переменная типа функция
	if replacer, err = makeReplacerFunction(britishAmerican); err != nil {
		return err
	}
	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	for !eof {

	}
	return nil
}

