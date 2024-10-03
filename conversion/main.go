package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	input, _ := read.ReadString('\n')
	value, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println("Error is ", error)
	} else {
		fmt.Println("Input is ", value)
	}
}
