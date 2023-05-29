package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://dmoj.ca/problem/16bitswonly

func main() {
	// get number of loops
	num_loops := get_num_loops()
	fmt.Println(num_loops)

	// get input
	for i := 0; i < num_loops; i++ {
		process_input()
	}

	// TO DO: check if CPUs are 16 bit only
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
func process_input() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

}
