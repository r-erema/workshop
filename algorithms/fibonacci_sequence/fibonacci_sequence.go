package fibonaccisequence

const fibPrev, fibNext = 1, 2

// FibonacciRecursive calculates Fibonacci number recursively.
// Time O(N^2), since we recount numbers in recursion
// Space O(N), since we use stretch recursion stack.
func FibonacciRecursive(number int) int {
	if number <= 1 {
		return number
	}

	return FibonacciRecursive(number-fibPrev) + FibonacciRecursive(number-fibNext)
}

// FibonacciCache calculates Fibonacci number using memoization.
// Time O(N), since we iterate input one time and get counted result from cache
// Space O(N), since we involve map as a cache which is equal to input length.
func FibonacciCache(number int) int {
	if number <= 1 {
		return number
	}

	return CacheHelper(number, map[int]int{0: 0, 1: 1})
}

func CacheHelper(number int, cache map[int]int) int {
	if _, ok := cache[number]; !ok {
		cache[number] = CacheHelper(number-fibPrev, cache) + CacheHelper(number-fibNext, cache)
	}

	return cache[number]
}

// FibonacciIterative calculates Fibonacci number iteratively.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func FibonacciIterative(number int) int {
	if number <= 1 {
		return number
	}

	prev1, prev2 := 0, 1
	for i := 2; i <= number; i++ {
		prev1, prev2 = prev2, prev1+prev2
	}

	return prev2
}
