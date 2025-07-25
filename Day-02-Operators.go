package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	// Write your code here
	tip := meal_cost * float64(tip_percent) / 100
	tax := meal_cost * float64(tax_percent) / 100

	total := meal_cost + tip + tax

	roundedTotal := int(total + 0.5)

	fmt.Printf("%d\n", roundedTotal)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	meal_cost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
