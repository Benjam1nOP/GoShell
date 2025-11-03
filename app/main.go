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

		command = strings.TrimRight(command, "\r\n")
		
		fmt.Printf("%s: command not found\n",command)
	}
}