package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputInt, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	inputFloat, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	inputString := scanner.Text()

	fmt.Println(i + uint64(inputInt))
	fmt.Printf("%.1f\n", d+inputFloat)
	fmt.Println(s + inputString)

}
