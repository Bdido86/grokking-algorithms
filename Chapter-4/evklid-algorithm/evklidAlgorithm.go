package main

import "fmt"

func main() {
	nod := findNOD(78, 294, 570, 36)
	fmt.Printf("nod=%d", nod)
}

// Поиск наибольший общий делитель(НОД)
func findNOD(nums ...int) int {
	if len(nums) == 1 {
		return abs(nums[0])
	}

	if len(nums) == 2 {
		return nod(nums[0], nums[1])
	}

	nums = append(nums[2:], nod(nums[0], nums[1]))
	return findNOD(nums...)
}

func nod(x int, y int) int {
	x = abs(x)
	y = abs(y)
	if x == y {
		return x
	}

	var val int
	if x > y {
		val = x % y
		if val == 0 {
			return y
		}
		val = nod(y, val)
	} else {
		val = y % x
		if val == 0 {
			return x
		}
		val = nod(x, val)
	}

	return val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
