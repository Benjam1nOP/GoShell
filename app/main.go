package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	existingfunc := []string {"echo","exit","type"} // array for the existing commands

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
		switch {
			case cmd == "exit 1":
				os.Exit(1)
			case cmd == "exit 0":
				os.Exit(0)

			case strings.HasPrefix(cmd, "type"):{
				arg := strings.TrimSpace(cmd[4:])
				if arg == ""{	// TrimSpace will Trim "type echo" to "echo" and for "type " (type and space) from " " to ""
					fmt.Fprintln(os.Stderr, "Invalid no. of arguments for Command Type")
					break
				}
				if slices.Contains(existingfunc, arg){ // checking if it is an existing command or not in the array
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arg)
				}else {
					fmt.Fprintf(os.Stderr, "%s: not found\n", arg)
				}
			}

			case strings.HasPrefix(cmd, "echo"):{
				arg := strings.TrimSpace(cmd[4:])
				//fmt.Println(cmd[4:])
				if arg == "" {
					fmt.Fprintln(os.Stderr, "Invalid no. of parameter for echo")
					break
				}
				fmt.Fprintln(os.Stdout, cmd[5:])
			}

			default: 
			fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
		}
		/*if strings.HasPrefix(cmd, "type"){
			arg := strings.TrimSpace(cmd[5:])
			if arg == "exit" || arg == "echo"{
				fmt.Printf("%s is a shell builtin", arg)
			}else{
			fmt.Printf("%s: not found\n", cmd)
			}
			return
		}

		if strings.HasPrefix(cmd, "echo"){ // checking prefix
			if len(cmd)==4{
				fmt.Println("Invalid no. of parameter for echo")
				os.Exit(1)
			}
			fmt.Println(cmd[5:])
			return
		}
		fmt.Printf("%s: command not found\n", command) */
	}
}