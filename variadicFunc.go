//variadic function

package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
