// URL: https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/best-index-1-45a2f8ff/description/
package main

import (
	"fmt"
)

var a = []int32{}
var l = 0

func main() {
	fmt.Scanf("%d", &l) // get the input value
	a = make([]int32, l)
	y := make([]interface{}, l)
	for i := range a {
		y[i] = &a[i]
	}
	fmt.Scanln(y...)
	//a = a[:n]

	var bestIndexSS int32 = 0
	for i := range a {
		il := l - i // offset by index
		n := 0      // nth term
		d := 0      // last sum
		for {
			ls := sum(n) // get the sum of the nth term
			if ls > il { // check if it is larger than the length
				break // exit out of the loop
			}
			d = ls
			n++
		}

		var ss int32 = 0 // variable for the sum
		ll := i + d      // length of the loop
		for i < ll {
			ss += a[i] // add to sum
			i++        // iterate
		}
		// check if the best
		if ss > bestIndexSS {
			// make a new best
			bestIndexSS = ss
		}
	}
	fmt.Println(bestIndexSS)
}
func sum(n int) int {
	return ((n * n) + n) / 2
}
