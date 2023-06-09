package main

import "fmt"

func Sum(set []int) int {
	var result int
	for _, num := range set {
		result += num
	}
	return result
}

func main() {
	fmt.Println("hello")
}
