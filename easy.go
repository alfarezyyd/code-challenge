package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	initNumber := 0
	copyOfParameter := x
	// 123
	for copyOfParameter != 0 {

		modulo := copyOfParameter % 10      // 3 2
		initNumber = initNumber*10 + modulo // 32
		copyOfParameter /= 10               // 12
	}
	fmt.Println(initNumber)
	return initNumber == x
}

// 121 % 10 = 1
// 10 20 10
