package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"net"
	"os"
	"strconv"
)

const(
	// Используемый tcp протокол
	tcpProtocol = "tcp4"

	// Длина генерируемого rsa ключа
	keySize = 1024

	// Максимальная длина шифруемого сообщения в байтах
	readWriterSize = keySize/8
)


var connectAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 4000}

func main() {
	// Соединяемся с сервером
	c := connectTo()

	// Буферизирует всё, что приходит от соединения "c"
	buf := bufio.NewReader(c)

	// Создаём приватный ключ в составе которого уже есть публичный ключ
	k, err := rsa.GenerateKey(rand.Reader, keySize)
	checkErr(err)

	// Отправляем серверу публичный ключ
	sendKey(c, k)

	// В цикле принимаем зашифрованные сообщения от сервера
	for {
		// Получаем зашифрованное сообщение в байтах
		cryptMsg := getBytes(buf, readWriterSize)

		// Расшифровываем сообщение
		msg, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, k, cryptMsg, nil)

		// Проверяем на ошибку
		checkErr(err)

		// Выводим расшифрованное сообщение
		fmt.Println(string(msg))
	}
}

func checkErr(err error){
	if err != nil {
		// Выводим текст ошибки
		fmt.Println(err)
		// Завершаем программу
		os.Exit(1)
	}
}

// Считываем с командной строки нужный нам порт и пытаемся соединится с сервером
func connectTo() *net.TCPConn {
	// Выводим текст "Enter port:" без перехода но новую строку
	fmt.Print("Enter port:")

	// Считываем число с консоли в десятичном формате "%d"
	fmt.Scanf("%d", &connectAddr.Port)

	fmt.Println("Connect to", connectAddr)

	// Создаём соединение с сервером
	c, err := net.DialTCP(tcpProtocol, nil, connectAddr)
	checkErr(err)
	return c
}

// Функция в определённом порядке отправляет PublicKey
func sendKey(c *net.TCPConn, k *rsa.PrivateKey) {
	// Говорим серверу что сейчас будет передан PublicKey
	c.Write([]byte("CONNECT\n"))
	// []byte() конвертирует "строку" в срез байт
	// передаём N типа *big.Int конвертированного в строку
	c.Write([]byte(k.PublicKey.N.String() + "\n"))
	// String() конвертирует *big.Int в string
	// передаём E типа int
	c.Write([]byte(strconv.Itoa(k.PublicKey.E) + "\n"))
	// strconv.Itoa() конвертирует int в string
}

// Читает и освобождает определённый кусок буфера
// Вернёт срез байт
func getBytes(buf *bufio.Reader, n int) []byte {
	// Читаем n байт
	bytes, err := buf.Peek(n)
	checkErr(err)
	// Освобождаем n байт
	skipBytes(buf, n)
	return bytes
}

// Освобождает, пропускает определённое количество байт
func skipBytes(buf *bufio.Reader, skipCount int) {
	for i := 0; i < skipCount; i++ {
		buf.ReadByte()
	}
}
