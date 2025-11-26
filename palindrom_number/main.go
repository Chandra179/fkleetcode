package main

import "fmt"

func main() {
	x := 123211
	num := x
	reversed := 0

	if num < 0 {
		fmt.Println("false")
	}
	for num != 0 {
		reversed = 10*reversed + num%10
		num /= 10
	}
	fmt.Println(reversed == x)
}
