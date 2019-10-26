package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime" // юзаем для определения типа ОС
)

type polar struct {
	radius float64
	angle float64
}

type cartesian struct {
	x float64
	y float64
}

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

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <- questions
			angle := polarCoord.angle * math.Pi / 180.0 // преобразование градусов в радианы
			x := polarCoord.radius * math.Cos(angle)
			y := polarCoord.radius * math.Sin(angle)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02f angle=%.02f° → Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian)  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n') //читает строку, пока не дойдет до \n
		if err != nil {
			break
		}
		var radius, angle float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &angle); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, angle}
		coord := <- answers
		fmt.Printf(result, radius, angle, coord.x, coord.y)
	}
	println()
}
