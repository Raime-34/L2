package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		var instructions = make([]IShite, 0)

		lines := strings.Split(text, " |")
		for _, subStr := range lines {
			subStr = strings.Trim(subStr, " ")
			if subStr == "" {
				continue
			}
			instructions = checkKWord(subStr, instructions)
		}

		for _, i := range instructions {
			i.doSomething()
		}

	}

}

func checkKWord(subStr string, instructions []IShite) []IShite {
	inputs := strings.Split(subStr, " ")
	switch inputs[0] {
	case "cd":
		return append(
			instructions,
			&CD{
				Instruction{
					input: inputs[1],
				},
			})
	case "pwd":
		return append(
			instructions,
			&PWD{},
		)
	case "echo":
		return append(
			instructions,
			&Echo{
				Instruction{
					input: inputs[1],
				},
			},
		)
	case "ps":
		return append(
			instructions,
			&PS{},
		)
	case "kill":
		return append(
			instructions,
			&Kill{
				Instruction{
					input: inputs[1],
				},
			},
		)
	case "fork":
		return append(
			instructions,
			&Fork{},
		)
	case "exec":
		return append(
			instructions,
			&Exec{
				Instruction{
					input: inputs[1],
				},
			},
		)
	case "netcat":
		return append(
			instructions,
			&Netcat{},
		)
	default:
		return instructions
	}
}

// #region IShite
type IShite interface {
	doSomething()
}

type Instruction struct {
	input  string
	result interface{}
}

func (i *Instruction) doSomething() {}

// #endregion

// #region cd
type CD struct {
	Instruction
}

func (i *CD) doSomething() {
	cd(i.input)
}

func cd(path string) {
	home, _ := os.UserHomeDir()
	err := os.Chdir(filepath.Join(home, path))

	if err != nil {
		err = os.Chdir(path)

		if err != nil {
			panic(err)
		}
	}

}

// #endregion

// #region pwd
type PWD struct {
	Instruction
}

func (i *PWD) doSomething() {
	fmt.Println(pwd())
}

func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

// #endregion

// #region echo
type Echo struct {
	Instruction
}

func (i *Echo) doSomething() {
	echo(i.input)
}

func echo(msg string) {
	fmt.Println(msg)
}

// #endregion

// #region ps
type PS struct {
	Instruction
}

func (i *PS) doSomething() {
	for name, pid := range ps() {
		fmt.Println(name, pid)
	}
}

func ps() map[string]int {
	processes, _ := process.Processes()
	m := make(map[string]int, len(processes))
	for _, process := range processes {
		name, err := process.Name()

		if err == nil {
			m[name] = int(process.Pid)
		}
	}
	return m
}

// #endregion

// #region kill
type Kill struct {
	Instruction
}

func (i *Kill) doSomething() {
	pid, err := strconv.Atoi(i.input)

	if err != nil {
		return
	}

	kill(pid)
}

func kill(pid int) error {
	process, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	err = process.Kill()

	return err
}

// #endregion

// #region fork
type Fork struct {
	Instruction
}

func (i *Fork) doSomething() {
	fork()
}

func fork() {
	cmd := exec.Command("cmd", "/c", "start", os.Args[0])
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// #endregion

// #region exec
type Exec struct {
	Instruction
}

func (i *Exec) doSomething() {
	myExec(i.input)
}

func myExec(exe string) {
	cmd := exec.Command(exe)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// #endregion

// #region netcat
type Netcat struct{}

func (n *Netcat) doSomething() {
	startClient(":8080")
}

func startClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}
}

// #endregion
