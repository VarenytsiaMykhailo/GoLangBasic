package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//слушать порт (Bind на порт ОС). Здесь может быть tcp/udp/tcp4/tcp6
	listener, err := net.Listen("tcp", ":5000")
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

		go func(conn net.Conn) {
			defer func() {
				fmt.Println("Connection closed")
				conn.Close()
			}()
			// создаем Reader для чтения информации из сокета
			//начнем читать из conn побайтово и выводить в консоль
			bufReader := bufio.NewReader(conn)
			fmt.Println("Start reading")
			for {
				//побайтово читаем
				rbyte, err := bufReader.ReadByte()
				if err != nil {
					fmt.Println("Can not read!", err)
					break
				}
				fmt.Print(string(rbyte))
			}
		}(conn)
	}
}
