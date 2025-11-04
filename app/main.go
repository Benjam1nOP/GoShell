package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Fprint(os.Stdout, "$ ")

		command, err:= reader.ReadString('\n')

		if err!=nil{
			if err == io.EOF{
				break
			}
			fmt.Printf("-Error Reading the input")
			os.Exit(1)
		}
		//fmt.Printf("%d",len(command))	check what is getting stored after input and it is text with \r\n

		//command = strings.TrimRight(command, "\r\n")
		if command == "exit 1\r\n" {
			os.Exit(1)
		}
		if command == "exit 0\r\n" {
			os.Exit(0)
		}
		fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
	}
}