package main

import "fmt"

func main() {
	score := []int{10, 9, 8, 7, 7, 7, 5, 5}
	score1 := []int{0, 0, 0, 0}
	fmt.Println(Finalist(score))
	fmt.Println(Finalist(score1))

}

func Finalist(n []int) int {
	var res int

	for i, v := range n {
		if i == len(n)-1 {
			break
		}
		if v == n[len(n)-1] {
			break
		}
		if v >= n[i+1] && n[i] != n[len(n)-1] {
			res++
		}
	}

	return res
}
