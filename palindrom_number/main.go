package main

import "fmt"

func main() {
	x := 121
	num := x
	reversed := 0

	if num < 0 {
		fmt.Println("false")
	}

	// reversed = 10*0 + 1 = 1
	// num  = 121 / 10 = 12

	// reversed = 10*1 + 2 = 12
	// num = 12 / 10 = 1

	// reversed = 10*12 + 1 = 121
	// num = 1 / 10 = 0
	for num != 0 {
		reversed = 10*reversed + num%10
		num /= 10
	}
	fmt.Println(reversed == x)
}
