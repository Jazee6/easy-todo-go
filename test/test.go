package main

import "fmt"

func main() {
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)
	count := 0
	for A := 1; A*A <= N; A++ {
		for B := 1; A*B <= N; B++ {
			C := N - A*B
			if C <= 0 {
				break
			}
			for D := 1; C*D <= N; D++ {
				if A*B+C*D == N {
					count++
				}
			}
		}
	}
	fmt.Printf("The number of quadruples of positive integers (A,B,C,D) such that AB+CD=%d is %d\n", N, count)
}
