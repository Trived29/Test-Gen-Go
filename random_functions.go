package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func reverseString(s string) string {
	return strings.Join(strings.Split(s, "")[::-1], "")
}


func factorial(n int) int {
	if n < 0 {
		return -1 // Error condition for negative input
	} else if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
