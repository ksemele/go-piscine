package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	//"go/token"
)

const (
	USAGE = "USAGE: pls use correctly!"
	ERROR = -1
	)
/*
** func check incoming args to valide args
*/

func ftCheckArgs(args []string) (res int) {
	validArgs := []string{"-mn", "-md", "-mo", "-sd"}
	for _, v := range args {
		res = -1
		for _, v_ := range validArgs {
			if v_ == v {
				res = 0
			}
		}
		if res == -1 {
			return
		}
	}
	return
}

/*
** scan lines from Stdio
** numbers - slice (dynamic array) with 0 capacity
** str - temporary string for Scan stdio
*/

func ftReadNumbs(numbers *[]int) {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str = scanner.Text()
		if s, err := strconv.ParseInt(str, 10, 32);
			err == nil && int(s) >= -100000 && int(s) <= 100000 {
			*numbers = append(*numbers, int(s))
		} else {
			fmt.Println("Invalid input: ", str)
			os.Exit(255)
		}
	}
	if scanner.Err() != nil {
		os.Exit(255)
	}

}

func main() {

/*
** read args for flags management
*/
	args := os.Args[1:]
	switch lenArgs := len(args); {
	case lenArgs > 0 && lenArgs <= 4:
		//fmt.Println("Founded ARGS: ", args)
		check := ftCheckArgs(args)
		//fmt.Println("\n\nargs ", check) //check args
		if check == -1 {
			fmt.Println(USAGE)
			os.Exit(ERROR)
		}
	default:
		fmt.Println("NO ARGS! set all flags to 1")
		args = append(args, "-mn", "-md", "-mo", "-sd")
		fmt.Println(args)
	}



	numbers := make([]int, 0)
	ptr := &numbers
	ftReadNumbs(ptr)

/*
** sorting slice numbers
*/
	sort.Ints(numbers)

	fmt.Printf("readed numbers:\n")
	fmt.Println(numbers)
	fmt.Println("End")
}
