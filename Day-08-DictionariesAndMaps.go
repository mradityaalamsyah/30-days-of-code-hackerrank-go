package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	phoneBook := make(map[string]string)

	for i := 0; i < n && scanner.Scan(); i++ {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 2 {
			name := parts[0]
			number := parts[1]
			phoneBook[name] = number
		}
	}

	for scanner.Scan() {
		query := strings.TrimSpace(scanner.Text())
		if val, found := phoneBook[query]; found {
			fmt.Printf("%s=%s\n", query, val)
		} else {
			fmt.Println("Not found")
		}
	}
}
