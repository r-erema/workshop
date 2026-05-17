package evaluatereversepolishnotation

import (
	"strconv"
)

// EvaluateReversePolishNotation evaluates reverse polish notation tokens.
// Time O(n), since we iterate each element of input
// Space O(n), since we involve stack equals input.
func EvaluateReversePolishNotation(tokens []string) int {
	var (
		stack              []int
		operand1, operand2 int
	)

	for _, token := range tokens {
		switch token {
		case "+":
			operand1, operand2, stack = GetAndPopLastOperand(stack)
			stack = append(stack, operand1+operand2)
		case "-":
			operand1, operand2, stack = GetAndPopLastOperand(stack)
			stack = append(stack, operand1-operand2)
		case "*":
			operand1, operand2, stack = GetAndPopLastOperand(stack)
			stack = append(stack, operand1*operand2)
		case "/":
			operand1, operand2, stack = GetAndPopLastOperand(stack)
			stack = append(stack, operand1/operand2)
		default:
			s, _ := strconv.Atoi(token)
			stack = append(stack, s)
		}
	}

	return stack[0]
}

// GetAndPopLastOperand pops the last two operands from the stack.
func GetAndPopLastOperand(stack []int) (int, int, []int) {
	return stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
}
