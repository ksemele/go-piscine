package main

import (
	"bufio"
	"fmt"
	"math"
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
func ftGetMean(numbers []int) (mean float32) {
	for _, v := range numbers {
		mean += float32(v)
	}
	return mean / float32(len(numbers))
}

func ftMean(numbers []int, args []string)  {
	for _, v_ := range args {
		if v_ == "-mn" {
			fmt.Printf("Mean: %.2f\n", ftGetMean(numbers))
		}
	}
}

func ftMedian(numbers []int, args []string)  {
	for _, v_ := range args {
		if v_ == "-md" {
			if len(numbers) % 2 != 0 {
				fmt.Printf("Median: %.2f\n", float32(numbers[len(numbers) / 2]))
			} else {
				fmt.Printf("Median: %.2f\n",
					(float32(numbers[(len(numbers) / 2) - 1]) + float32(numbers[len(numbers) / 2])) / 2)
			}
		}
	}
}

/*
** func find mode (often occurance in []int)
** Create a map and populated it with each value in the slice
** mapped to the number of times it occurs
** Find the smallest item with greatest number of occurance in
** the input slice
*/

func ftGetMode(array []int) (mode int) {
	countMap := make(map[int]int)
	for _, value := range array {
		countMap[value] += 1
	}
	max := 0
	for _, key := range array {
		freq := countMap[key]
		if freq > max {
			mode = key
			max = freq
		}
	}
	return
}

func ftMode(numbers []int, args []string) {
	for _, v_ := range args {
		if v_ == "-md" {
			fmt.Printf("Mode: %.2f\n", float32(ftGetMode(numbers)))
		}
	}
}

func ftSdeviation(numbers []int, args []string) {
	for _, v_ := range args {
		if v_ == "-sd" {
			mean := ftGetMean(numbers)
			var res float32
			for _, v := range numbers {
				res += (float32(v) - mean) * (float32(v) - mean)
			}
			fmt.Printf("SD: %.2f\n", math.Sqrt(float64(res / float32(len(numbers)))))
		}
	}
}

func main() {

/*
** read args for flags management
*/
	args := os.Args[1:]
	switch lenArgs := len(args); {
	case lenArgs > 0 && lenArgs <= 4:
		check := ftCheckArgs(args)
		if check == -1 {
			fmt.Println(USAGE)
			os.Exit(ERROR)
		}
	default:
		//fmt.Println("NO ARGS! set all flags to 1")//todo del
		args = append(args, "-mn", "-md", "-mo", "-sd")
		//fmt.Println(args)//todo del
	}



	numbers := make([]int, 0)
	ptr := &numbers
	ftReadNumbs(ptr)

	fmt.Printf("readed numbers:\n")//todo del
	fmt.Println(numbers)//todo del
/*
** sorting slice numbers
*/
	sort.Ints(numbers)

	ftMean(numbers, args)
	ftMedian(numbers, args)
	ftMode(numbers, args)
	ftSdeviation(numbers, args)
	//fmt.Println("End")//todo del
}
