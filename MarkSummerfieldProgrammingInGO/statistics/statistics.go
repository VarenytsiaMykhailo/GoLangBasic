package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	pageTop = `
	<!DOCTYPE HTML><html>
	<head>
		<style>.error{color:#FF0000;}</style>
	</head>
	<title>Statistics</title>
	<body>
		<h3>Statistics</h3>
		<p>Computes basic statistics for a given list of numbers</p>
	`
	form = `
	<form action="/" method="POST">
	<label for="numbers">Numbers (comma or space-separated):</label><br />
	<input type="text" name="numbers" size="30"><br />
	<input type="submit" value="Calculate">
	</form>
	`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

/*
Функция http.HandleFunc() принимает два аргумента: путь и ссылку на функцию, которую
следует вызвать при поступлении запроса по указанному пути.
Функция должна иметь сигнатуру func(http.ResponseWriter, *http.Request).
Значение типа http.ResponseWriter - ответ клиенту от сервера (его мы должны сформировать в функции)
Значение типа *http.Request - запрос, пришедший от клиента. По нему мы можем определить, что хочет клиент и какой ответ ему сформировать.

Функция http.ListenAndServe(":9001", nil):
принимает первым аргументом адрес. Когда указывается только номер порта, автоматически предполагается, что адрес соответствует локальному компьютеру – с тем же
успехом можно было бы использовать адрес "localhost:9001" или
"127.0.0.1:9001". (Номер порта для данного приложения был выбран совершенно произвольно, вместо него можно указать другой,
если использование этого номера вызывает конфликты с существующим сервером.)
Вторым аргументом принимается тип сервера. Обычно в нем передается nil, чтобы выбрать тип по умолчанию
*/
func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

/*
В аргументе writer передается значение, куда должен записываться ответ (в формате HTML), а в аргументе
request – подробная информация о запросе.
*/
func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() //Анализ аргументов. Должен вызываться перед записью в ответ
	fmt.Println("r.URL", request.URL)
	fmt.Println("r.Body", request.Body)
	fmt.Println("r.GetBody", request.GetBody)
	fmt.Println("r.Header", request.Header)
	fmt.Println("r.Host", request.Host)
	fmt.Println("r.Method", request.Method)
	fmt.Println("r.Proto", request.Proto)
	fmt.Println("r.RemoteAddr", request.RemoteAddr)

	fmt.Println(request.Form) // вывод информации о форме на стороне сервера
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	fmt.Println("FORM:")
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false //при первом отображении данные отсутствуют
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers) // аргумент numbers функции getStats тоже отсортируется т.к. numbers и stats.numbers ссылаются в памяти на один массив
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return stats
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}
