package main

import "fmt"

const spb = 3         // seats per bench
const bps = 4         // benches per section
const spc = spb * bps // seats per section

var setTypes = [3]string{"WS", "MS", "AS"}
var st = ""

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for j := 0; j < n; j++ {
		var i int
		// get input data
		fmt.Scanf("%d", &i)

		// get section seat no and section no
		ssn, sNo := getRemainder(i-1, spc)

		// get bench seat no. and section bench no
		bsn, sbn := getRemainder(ssn, 3)

		if sbn == 1 || sbn == 3 {
			st = setTypes[2-bsn]
		} else {
			st = setTypes[bsn]
		}

		opSet := spc - ssn + (sNo * spc)

		fmt.Printf("%v %v\n", opSet, st)
	}
}

func getRemainder(n, d int) (r, q int) {
	r = n % d
	q = n / d
	return r, q
}