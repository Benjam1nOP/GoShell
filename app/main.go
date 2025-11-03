package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	fmt.Fprint(os.Stdout, "$ ")
	command, err:= bufio.NewReader(os.Stdin).ReadString('\n')
	if err!=nil{
		fmt.Printf("-Error Reading the input")
		os.Exit(1)
	}
	fmt.Printf("%s : command not found ",command)
}