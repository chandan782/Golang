// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func increment(i int) func() int {
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	i := 5
	f := increment(i)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
