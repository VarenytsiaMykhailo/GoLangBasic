package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	who := "World!" //будет выводиться world, если в консоли ничего не передать
	if len(os.Args) > 1 { //os.Args[0] - вызов программы в терминале: hello или hello.exe
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
