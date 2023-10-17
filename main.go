package main

import "fmt"

func main() {
	limit := 20                     // The upper limit for the numbers being considered
	factors := []int{3, 5}          // The factors that are used to find the multiples
	sum := 0                        // Variable to keep track of the sum of unique multiples
	multiples := make(map[int]bool) // A map to keep track of the multiples

	// Loop through each factor
	for _, factor := range factors {
		// Iterate over multiples of the factor less than the limit
		for i := factor; i < limit; i += factor {
			// Check if the multiple hasn't been encountered before
			if !multiples[i] {
				// If it's a new multiple, mark it as encountered in the map and add it to the sum
				multiples[i] = true
				sum += i
			}
		}
	}

	// Print the final sum of unique multiples
	fmt.Println("The sum of unique multiples is:", sum)
}
