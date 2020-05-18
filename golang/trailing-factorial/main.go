package main

import (
	"fmt"
	"math/big"
)

func trailingZeroes(n int) int {
	fives := 0
	for n >= 5 {
		fives += n / 5
		n = n / 5
	}
	return fives
}

func main() {
	fmt.Println(trailingZeroes(5))
}

func trailingZeroesSimple(n int) int {
	if n == 0 {
		return 0
	}

	factorial := factorial(n)

	zeroes := 0
	strFact := factorial.String()
	for i := len(strFact) - 1; i > 0; i-- {
		if strFact[i] != '0' {
			break
		}
		zeroes++
	}

	return zeroes
}

func factorial(n int) *big.Int {
	factorial := big.NewInt(1)
	for n != 1 {
		factorial.Set(factorial.Mul(big.NewInt(int64(n)), factorial))
		n--
	}
	return factorial
}

// func trailingZeroes(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	factorial := factorial(n)
// 	zeroes := 0
// 	if factorial%10 == 0 {
// 		for factorial%10 == 0 {
// 			zeroes++
// 			factorial = factorial / 10
// 		}
// 	}

// 	return zeroes
// }

// func factorial(n int) int {
// 	factorial := 1
// 	for n != 1 {
// 		factorial = factorial * n
// 		n--
// 	}
// 	return factorial
// }
