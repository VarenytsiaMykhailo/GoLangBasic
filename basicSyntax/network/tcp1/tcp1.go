package main

import (
	"encoding/gob"
	"fmt"
	"net"
)
/*
Этот пример использует пакет encoding/gob,
который позволяет легко кодировать выходные данные,
чтобы другие программы на Go (или конкретно эта программа, в нашем случае) могли их прочитать.
Дополнительные способы кодирования доступны в пакете encoding (например encoding/json),
а так-же в пакетах сторонних разработчиков (например, можно использовать labix.org/v2/mgo/bson для работы с BSON).
*/
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

func server()  {
	//слушать порт (Bind на порт ОС). Здесь может быть tcp/udp/tcp4/tcp6
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//принятие соеднинения (ждем, пока придет клиент)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Can not connect!", err)
			conn.Close()
			continue
		}
		fmt.Println("Connected")
		//обработка соединения
		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn)  {
	//получение сообщения
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	conn.Close()
}

func client() {
	//соединиться с сервером
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//послать сообщение
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}