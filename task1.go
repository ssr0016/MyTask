package main

import "fmt"

func main() {
	limit := 20
	factors := []int{3, 5}
	sum := 0
	multiples := make(map[int]bool)

	for _, factor := range factors {
		for i := factor; i < limit; i += factor {
			if !multiples[i] {
				multiples[i] = true
				sum += i
			}
		}
	}

	fmt.Println("The sum of unique multiples is:", sum)
}
