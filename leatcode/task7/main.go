package main

import (
	"fmt"
)

func main() {
	digits := []int{9, 9}

	fmt.Println(PlusOne(digits))
}

func PlusOne(digits []int) []int {
	var increment int

	for i := len(digits) - 1; i >= 0; i-- {
		if i != 0 && digits[i] == 9 {
			digits[i] = 0
			continue
		}
		if i == 0 && digits[i] == 9 {
			digits[i] = 1
			increment++
			break
		}
		digits[i] += 1
		break
	}

	if digits[0] == 1 && increment != 0 {
		digits = append(digits, 0)
		return digits
	}

	return digits
}
