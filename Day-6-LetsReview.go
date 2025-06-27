package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	fmt.Scan(&T)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i <= T; i++ {
		scanner.Scan()
		text := scanner.Text()

		even := ""
		odd := ""

		for S, char := range text {
			if S%2 == 0 {
				even += string(char)
			} else {
				odd += string(char)
			}

		}
		if text != "" {
			fmt.Println(even, odd)
		}
	}
}
