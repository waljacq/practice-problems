package main

import "fmt"

func change(amount int, coins []int) int {

	dp := make([]int, amount+1)
	dp[0] = 1
	// if amount == 0 {
	//     return 1
	// }

	// if amount < 0 || len(coins) == 0 {
	//     return 0
	// }

	for i := 0; i < len(coins); i++ {
		for j := 1; j < len(dp); j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
			// change(amount - coins[len(coins)-1], coins)
			// change(amount, coins[:len(coins)-1])
		}
	}

	return dp[amount]
	// return change(amount - coins[len(coins)-1], coins) + change(amount, coins[:len(coins)-1])
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
