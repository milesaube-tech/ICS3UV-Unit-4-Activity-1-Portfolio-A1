// Author: Miles Aube
// Version: 1.0.0
// Date: 2025-12-05
// Fileoverview: This program prints out the times table.

package main

import (
	"fmt"
)

func main() {

	// set variables
	var counter int = 0
	var output string = ""

	// for loop
	for counter = 100; counter > 0; counter = counter - 5 {
		output = output + fmt.Sprintf("%d,", counter)
	}

	// add final 0
	output = output + "0"

	fmt.Println(output)
}