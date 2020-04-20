// URL: https://www.hackerearth.com/practice/basic-programming/input-output/basics-of-input-output/practice-problems/algorithm/best-index-1-45a2f8ff/description/
package main

import (
	"fmt"
)

var a = []int{}
var l = 0

func main() {
	fmt.Scanf("%d", &l) // get the input value
	a = make([]int, l)
	y := make([]interface{}, l)
	for i := range a {
		y[i] = &a[i]
	}
	n, _ := fmt.Scanln(y...)
	a = a[:n]

	bestIndexSS := 0
	for i := range a {
		d := calLastIndex(i)      // calculate the last index to calculate
		ss := calSpecialSum(i, d) // calculate special sum
		// check if the best
		if ss > bestIndexSS {
			// make a new best
			bestIndexSS = ss
		}
	}
	fmt.Println(bestIndexSS)
}

func calSpecialSum(i, d int) int {
	sum := 0   // variable for the sum
	l := i + d // length of the loop
	for i < l {
		sum += a[i] // add to sum
		i++         // iterate
	}
	return sum
}

// calculate the the last index for the special sum. since the index increases
// in a sequrence 1,3,6,10... which is the sum 1st nth term of the following sequence
//  1,2,3,4,5,6,7,8,9,10...
// so by calculating the last sum which is smaller than the length, we can calcuate the
// range of index to add.
func calLastIndex(i int) int {
	il := l - i // offset by index
	n := 0      // nth term
	s := 0      // last sum
	for {
		ls := sum(n) // get the sum of the nth term
		if ls > il { // check if it is larger than the length
			break // exit out of the loop
		}
		s = ls
		n++
	}
	return s
}

// calculate the sum of 1st n terms of following number sequence
//  1,2,3,4,5,6,7,8,9,10...
func sum(n int) int {
	return ((n * n) + n) / 2
}
