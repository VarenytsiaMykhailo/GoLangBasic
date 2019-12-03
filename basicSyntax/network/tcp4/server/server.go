package main
/*
Что делает приложение
[к] — клиент
[c] — сервер
1. По установленному TCP соединению, [к] передает публичный ключ rsa.
2. При помощи принятого публичного ключа, [c] шифрует и отправляет сообщения [к]
3. [к] расшифровывает и выводит сообщения.
*/

import (
	"bufio"
	// Пакет для кроссплатформенной генерации случайных чисел
	"crypto/rand"
	// При помощи этого пакета будем шифровать и дешифровать передаваемую информацию + там содержится тип данных rsa-ключа
	"crypto/rsa"
	// Для создания хешей методом sha1
	"crypto/sha1"
	"fmt"
	// Для работы с большими числами
	"math/big"
	// Пакет для передачи информации по Unix networks sockets, including TCP/IP, UDP протоколам. В данном случае будем использовать TCP протокол.
	"net"
	// Пакет для кроссплатформенной взаимодействия с операционной системой
	"os"
	// Для конвертации строковых данных в основные типы данных и обратно
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

/*Для того чтоб держать вместе соединение сервера и ключ от этого соединения «pubK»
объявим тип данных remoteConn как структуру:*/
type remoteConn struct {
	c *net.TCPConn
	pubK *rsa.PublicKey //указатель на ключ (на структуру типа rsa.PublicKey). Нужен для зашифровки передаваемых сообщений.
}

func main() {
	listen()
}

/*
Следующая функция объединяет, соединение и публичный ключ для шифрования этого соединения в структуру remoteConn.
Причём возвращает ссылку на remoteConn а не значение.
waitPubKey() — ожидает от «клиента» когда тот в определённой последовательности передаст PublicKey
*/
func getRemoteConn(c *net.TCPConn) *remoteConn{
	return &remoteConn{c: c, pubK: waitPubKey(bufio.NewReader(c))}
}

/*
В целях ознакомления будем обрабатывать возникающие ошибки таким образом:
Функция принимает одно значение err у которого тип error.
*/
func checkErr(err error){
	if err != nil {
		// Выводим текст ошибки
		fmt.Println("ERROR: ", err)
		// Завершаем программу
		os.Exit(1)
	}
}

/*
Функция принимает ссылку на буфер (*bufio.Reader)
который в свою очередь содержит все байты пришедшие от соединение «c».
Вернёт ссылку на структуру данных rsa.PublicKey
*/
func waitPubKey(buf *bufio.Reader) (*rsa.PublicKey) {

	// Читаем строку из буфера
	line, _, err := buf.ReadLine(); checkErr(err)

	// Так как тип line - []byte (срез байт)
	// то для удобства сравнения переконвертируем line в строку
	if string(line) == "CONNECT" {
		// Далее мы будем читать буфер в том же порядке, в котором приходят данные с клиента
		line, _, err := buf.ReadLine(); checkErr(err) // Читаем PublicKey.N

		// Создаём пустой rsa.PublicKey
		pubKey := rsa.PublicKey{N: big.NewInt(0)}
		// pubKey.N == 0
		// тип pubKey.N big.Int http://golang.org/pkg/big/#Int

		// Конвертируем полученную строку в big Int и запихиваем в pubKey.N big.Int
		pubKey.N.SetString(string(line), 10)
		// Метод SetString() получает 2 параметра:
		// string(line) - конвертирует полученные байты в строку
		// 10 - система исчисления используемая в данной строке
		// (2 двоичная, 8 восьмеричная, 10 десятичная, 16 шестнадцатеричная ...)

		// Читаем из буфера второе число для pubKey.E
		line, _, err = buf.ReadLine(); checkErr(err)

		// Используемый пакет strconv для конвертации тип string в тип int
		pubKey.E, err = strconv.Atoi(string(line)); checkErr(err)

		// возвращаем ссылку на rsa.PublicKey
		return &pubKey

	} else {
		// В этом случае дальнейшее действия программы не предусмотренно. Поэтому:
		// Выводим что получили
		fmt.Println("Error: unkown command from client", string(line))
		os.Exit(1) // Завершаем программу
		return nil
	}
}

/*
Следующая функция является методом для ссылки на переменную типа remoteConn
Проделывает ряд действий для зашифровки и отправки сообщения
*/
func (rConn *remoteConn) sendCommand(comm string) {

	// Зашифровываем сообщение
	eComm, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, rConn.pubK, []byte(comm), nil)
	// sha1.New() вернёт данные типа hash.Hash
	// С таким же успехм можно использовать sha512.New() sha256.New() ...
	// rand.Reader тип которого io.Reader позволяет не задумываясь о платформе генерировать
	// случайные числа из /dev/unrandom будь то Linux или CryptGenRandom API будь то Windows
	// rConn.pubK - публичный ключ который мы получили в func waitPubKey
	// []byte(comm) - конвертируем строку comm в срез байт ([]byte)
	checkErr(err) // проверяем на ошибки
	// Передаём зашифрованное сообщение клиенту по заранее установленному соединению
	rConn.c.Write(eComm)
	// rConn.c какого типа? - net.TCPConn у которого есть метод Write()
	// http://golang.org/pkg/net/#TCPConn.Write
}

/*
Объявим глобальную переменную listenAddr которая будет ссылкой на структуру типа net.TCPAddr
«Port: 0» в данном случае означает — любой свободный порт.
*/
var listenAddr = &net.TCPAddr{IP: net.IPv4(127,0,0,1), Port: 0}

/*
Ниже функция, которая оперирует ранее объявленными функциями и в конечном итоге
отправляет «клиенту» название сервера и приветствия на разных языках.
*/
func listen() {
	// Слушаем любой свободны порт
	l, err := net.ListenTCP(tcpProtocol, listenAddr); checkErr(err)

	// Выведем прослушиваемый порт
	fmt.Println("Listen port: ", l.Addr().(*net.TCPAddr).Port)
	// l == *net.TCPListener == ссылка на тип данных
	// .Addr() http://golang.org/pkg/net/#TCPListener.Addr == метод для *net.TCPListener который возвращает "интерфейс"
	// net.Addr http://golang.org/pkg/net/#Addr который в свою очередь содержит ссылку на TCPAddr - *net.TCPAddr
	// и два метода Network() и String()

	c, err := l.AcceptTCP(); checkErr(err)
	// На этом этапе программа приостанавливает свою работу ожидая соединения по прослушиваемому порту
	// AcceptTCP() - метод для *net.TCPListener http://golang.org/pkg/net/#TCPListener.AcceptTCP
	//Возвращает установленное соединение и ошибку

	fmt.Println("Connect from:", c.RemoteAddr())
	// Вот 3 варианта которые подставив в fmt.Print[f|ln]() получим одинаковый результат
	// 1. c.RemoteAddr()
	// 2. c.RemoteAddr().(*net.TCPAddr)
	// 3. c.RemoteAddr().String()
	// В первый двух случаях функции: fmt.Println(), fmt.Print(), fmt.Printf() попытаются найти метод String()
	// Иначе вывод будет таким как есть

	// Таким образом мы получим соединение и ключ которым можно зашифровать это соединение
	rConn := getRemoteConn(c)

	// Шифруем и отправляем сообщения
	rConn.sendCommand("Go Language Server v0.1 for learning")
	rConn.sendCommand("Привет!")
	rConn.sendCommand("Привіт!")
	rConn.sendCommand("Прывітанне!")
	rConn.sendCommand("Hello!")
	rConn.sendCommand("Salut!")
	rConn.sendCommand("ハイ!")
	rConn.sendCommand("您好!")
	rConn.sendCommand("안녕!")
	rConn.sendCommand("Hej!")
}