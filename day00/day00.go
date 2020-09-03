package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//arg := os.Args[0]

	//fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	//fmt.Println(arg)
	if len(argsWithoutProg) == 0 {
		fmt.Println("NO ARGS! :(")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		// handle error.
	}
}
