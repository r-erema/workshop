package generateparentheses

import (
	"strings"
)

// GenerateParenthesis generates all valid parentheses combinations.
// Time O(4^n/sqrt(n)), where n is the number of pairs of parentheses,
// in simple terms, the time complexity will be the nth Catalan number
//
// O(4^n/sqrt(n)) will be space complexity,
// where n is the number of pairs of parentheses,
// because of the recursion stack.
func GenerateParenthesis(count int) []string {
	var (
		stack, res []string
		backtrack  func(opened, closed int)
	)

	backtrack = func(opened, closed int) {
		if opened == closed && closed == count {
			res = append(res, strings.Join(stack, ""))

			return
		}

		if opened < count {
			stack = append(stack, "(")

			backtrack(opened+1, closed)

			stack = stack[:len(stack)-1]
		}

		if closed < opened {
			stack = append(stack, ")")

			backtrack(opened, closed+1)

			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)

	return res
}
