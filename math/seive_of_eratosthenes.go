package numbertheory

import "fmt"

func displaySeive(n int, prime []bool) {
	for i := 2; i <= n; i++ {
		if !prime[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func Seive(n int) {
	prime := make([]bool, n+1)

	for i := 2; i*i < n; i++ {
		if !prime[i] {
			for j := 2 * i; j <= n; j = j + i {
				prime[j] = true
			}
		}
	}

	displaySeive(n, prime)
}
