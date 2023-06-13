package numbertheory

func IsPrimeBasic(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime(n int) bool {
	flag := true

	if (n == 0) || (n == 1) {
		flag = false
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			flag = false
		}
	}

	return flag
}
