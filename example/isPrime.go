package main

import "fmt"

func main() {
	num := 10
	fmt.Println(isPrime(num))
}

func isPrime(n int) int64 {
	ct := int64(0)
	for i := 2; i < n; i++ {
		flag := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			ct++
		}
	}
	return ct
}
