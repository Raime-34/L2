package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	getPage()
}

func getPage() error {
	resp, err := http.Get("https://go.dev/tour/methods/21")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	buff := make([]byte, 1024)
	for {
		_, err := reader.Read(buff)
		fmt.Print(string(buff))
		buff = make([]byte, 1024)

		if err == io.EOF {
			break
		}
	}

	return nil
}
