package main

import (
	"fmt"
	"testing"
)

func Test_makeRequest(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(makeRequest())
	}
}
