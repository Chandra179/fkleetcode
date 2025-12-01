package main

import "fmt"

func main() {
	height := []int{6, 7, 4, 9, 1, 3, 1}
	i := 0
	j := len(height) - 1
	res := 0

	for i < j {
		minH := height[i]
		if height[j] < minH {
			minH = height[j]
		}
		area := (j - i) * minH
		if area > res {
			res = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	fmt.Println(res)
}
