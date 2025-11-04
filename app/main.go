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
		command = strings.TrimSpace(command)  // In linux env \n is sent after text in windows \r\n is sent
		if command == "exit 1" {
			os.Exit(1)
		}
		if command == "exit 0" {
			os.Exit(0)
		}
		fmt.Printf("%s: command not found\n", command)
	}
}