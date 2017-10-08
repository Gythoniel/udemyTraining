/*
Project Euler 205: Peter has nine four-sided (pyramidal) dice, each with faces numbered 1, 2, 3, 4.
Colin has six six-sided (cubic) dice, each with faces numbered 1, 2, 3, 4, 5, 6.

Peter and Colin roll their dice and compare totals: the highest total wins.
The result is a draw if the totals are equal.

What is the probability that Pyramidal Pete beats Cubic Colin?
Give your answer rounded to seven decimal places in the form 0.abcdefg
*/

package main

import (
	"fmt"
	"math"
)

// inspired by https://blog.dreamshire.com/project-euler-205-solution/
func main() {

	// p and c are the possibilities of reaching a certain score p(9-36) c(6-36)
	p := []float64{0, 0, 0, 1, 9, 45, 165, 486, 1206, 2598, 4950, 8451, 13051, 18351, 23607,
		27876, 30276, 30276, 27876, 23607, 18351, 13051, 8451, 4950, 2598, 1206, 486, 165, 45, 9, 1}
	c := []float64{1, 6, 21, 56, 126, 252, 456, 756, 1161, 1666, 2247, 2856, 3431, 3906, 4221,
		4332, 4221, 3906, 3431, 2856, 2247, 1666, 1161, 756, 456, 252, 126, 56, 21, 6, 1}
	var w float64
	var t float64 = math.Pow(4, 9) * math.Pow(6, 6)

	// range over c and multiply the amount of possibilities to arrive at that point in c
	// by the amount of possibilities that P has a higher score than this point in c

	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(p); j++ {
			w += p[j] * c[i]
		}
	}

	//  possibilities that p wins (w) divided by the total
	//  number of possibilities (t) results in P(win)

	fmt.Print("solution to euler 205: ", round(w/t, .5, 7))
}

// found at https://gist.github.com/DavidVaini/10308388
func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
