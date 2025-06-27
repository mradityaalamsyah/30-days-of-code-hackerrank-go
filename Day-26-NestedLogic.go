package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	day1, month1, year1 := int(0), int(0), int(0)
	day2, month2, year2 := int(0), int(0), int(0)

	fmt.Scan(&day1, &month1, &year1)
	fmt.Scan(&day2, &month2, &year2)

	fine := 0

	if year1 > year2 {
		fine = 10000
	} else if year1 == year2 {
		if month1 > month2 {
			fine = 500 * (month1 - month2)
		} else if month1 == month2 && day1 > day2 {
			fine = 15 * (day1 - day2)
		}
	}
	fmt.Println(fine)
}
