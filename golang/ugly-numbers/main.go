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

func ugly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}

	dp := make([]bool, (num/2)+1)

	uglies := []int{2, 3, 5}
	for i := 1; i < len(dp); i++ {
		for _, uglyNum := range uglies {
			if num%i == 0 {
				dp[i] = i%uglyNum == 0 || dp[i]
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

func findUgly(n int) int {
	return 0
}
