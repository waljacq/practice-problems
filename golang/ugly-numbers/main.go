package main

import "fmt"

func main() {
	num := 8

	if ugly(num) {
		fmt.Printf("%d is an ugly number\n", num)
	} else {
		fmt.Printf("%d is not an ugly number\n", num)
	}
}

// dp is incorrect
func ugly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 || num == 2 || num == 3 || num == 5 {
		return true
	}

	dp := make([]bool, num)

	uglies := []int{2, 3, 5}
	for i := 2; i < len(dp); i++ {
		for _, uglyNum := range uglies {
			if num%i == 0 {
				dp[i] = i%uglyNum == 0
			}

		}
	}
	return dp[(num / 2)]
}

func uglyRecursive(num int) bool {
	if num == 1 || num == 2 || num == 3 || num == 5 {
		return true
	}

	if num%2 == 0 {
		return ugly(num / 2)
	} else if num%3 == 0 {
		return ugly(num / 3)
	} else if num%5 == 0 {
		return ugly(num / 5)
	} else {
		return false
	}
}

func uglyIterative(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}

	// for i := 7; i < num+1; i++ {
	// 	if num%i == 0 {
	// 		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
	// 			return false
	// 		}
	// 	}
	// }

	if num%2 == 0 || num%3 == 0 || num%5 == 0 {
		for num%2 == 0 {
			num = (num / 2)
		}
		for num%3 == 0 {
			num = (num / 3)
		}
		for num%5 == 0 {
			num = (num / 5)
		}
	} else {
		return false
	}

	if num%2 != 0 && num%3 != 0 && num%5 != 0 && num != 1 {
		return false
	}
	return true
}

func findUgly(n int) int {
	return 0
}
