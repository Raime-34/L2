package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

//"github.com/reiver/go-telnet"

func main() {
	myTelnet()
}

func myTelnet() error {
	conn, err := net.DialTimeout("tcp", ":5555", 10*time.Second)

	if err != nil {
		fmt.Println(err)
		return err
	}

	go func() {
		for {
			msg, _ := bufio.NewReader(conn).ReadString('\n')
			if msg == "" {
				conn.Close()
				os.Exit(0)
			}
			fmt.Println(msg)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		if strings.Contains(line, string([]byte{4})) {
			conn.Close()
			return nil
		}
		conn.Write([]byte(line))
	}
}
