// https://dmoj.ca/problem/16bitswonly

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// get number of loops
	num_loops := get_num_loops()

	// get input
	var input []string
	for i := 0; i < num_loops; i++ {
		input = append(input, process_input())
	}

	// check if CPUs are 16 bit only
	for _, value := range input {
		fmt.Println(check_processor(value))
	}
}

// check if processor calculations are correct
func check_multiplition(a int, b int, p int) bool {
	if a*b == p {
		return true
	} else {
		return false
	}
}

// get number of loops from user input
func get_num_loops() int {
	var num_loops int
	fmt.Scanln(&num_loops)
	return num_loops
}

// process calculations from input
func process_input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// check each processor
func check_processor(calc string) string {
	s := strings.Fields(calc)
	num1, _ := strconv.Atoi(s[0])
	num2, _ := strconv.Atoi(s[1])
	result, _ := strconv.Atoi(s[2])
	if num1*num2 != result {
		return "16 BIT S/W ONLY"
	} else {
		return "POSSIBLE DOUBLE SIGMA"
	}
}
