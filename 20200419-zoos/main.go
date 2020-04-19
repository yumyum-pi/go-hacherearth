// url: https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/is-zoo-f6f309e7/
package main

import "fmt"

const maxStringSize = 20

func main() {

	var input string
	// get input data
	fmt.Scanf("%s", &input)

	c := 0
	l := len(input)
	// loop through all the strings
	for i := 0; i < l; i++ {
		if input[i] == 'z' {
			c++
			continue
		}
		break
	}
	x := 2 * c
	y := l - c
	if x == y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
