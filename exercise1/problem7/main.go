package main

import (
	"fmt"
	"strconv"
)

// highestDigit takes an integer as input and returns the highest digit in that number
func highestDigit(n int) int {
	// Convert the number to a string to access each digit
	numStr := strconv.Itoa(n)
	highest := 0

	// Iterate through each character in the string representation of the number
	for _, char := range numStr {
		// Convert the character to an integer
		digit := int(char - '0')
		// Update highest if the current digit is greater
		if digit > highest {
			highest = digit
		}
	}