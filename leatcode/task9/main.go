package main

import "fmt"

func main() {

	fmt.Println(MySqrt(256))

}

func MySqrt(x int) int {
	numsLow := 2
	numsHigh := 3

	if x <= 0 {
		return x
	}

	for true {
		if numsLow*numsLow == x {
			return numsLow
		}

		if numsLow*numsLow < x && numsHigh*numsHigh > x {
			return numsLow
		}
		numsLow = numsHigh
		numsHigh++
	}

	return 0
}
