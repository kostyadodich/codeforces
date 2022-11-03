package main

import "fmt"

func main() {
	score := []int{10, 9, 8, 7, 7, 7, 5, 5}
	score1 := []int{0, 0, 0, 0}
	fmt.Println(Finalist(score))
	fmt.Println(Finalist(score1))

}

func Finalist(n []int) int {
	low := n[len(n)-1]
	countLow := 0

	for i := len(n) - 1; i >= 0; i-- {
		if n[i] == low {
			countLow++
		} else {
			break
		}
	}

	return len(n) - countLow
}
