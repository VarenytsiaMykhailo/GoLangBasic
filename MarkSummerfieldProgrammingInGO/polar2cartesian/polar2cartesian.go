package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime" // юзаем для определения типа ОС
)

var prompt = "Enter a radius and an angle (in degrees), e.g.: 12.5 90  , or %s to quit."
func init()  {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "ctrl+Z, Enter")
	} else { // другая система (unix-подобная)
		prompt = fmt.Sprintf(prompt, "ctrl+D")
	}
}

func main()  {
	questions := make(chan polar)
	defer close(questions) //канал нужно закрыть
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)

}

type polar struct {
	radius float64
	angle float64
}

type cartesian struct {
	x float64
	y float64
}
