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
		cmd :=strings.ToLower(command) // to lower ECho to echo
		switch cmd{
			case "exit 1":
				os.Exit(1)
			case "exit 0":
				os.Exit(0)
		}

		if strings.HasPrefix(cmd, "echo"){
			if len(cmd)==4{
				fmt.Println("Invalid no. of parameter for echo")
				os.Exit(1)
			}
			fmt.Println(cmd[5:])
		}else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}