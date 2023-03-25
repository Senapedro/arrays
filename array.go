package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 5.3
	x[1] = 6
	x[2] = 9
	x[3] = 4.5
	x[4] = 7.5

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))
}
