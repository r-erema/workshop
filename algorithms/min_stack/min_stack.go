package minstack

// MinStack is a stack that supports retrieving the minimum element in constant time.
// Time O(1), we get all numbers from top of stack,
// even index of the min value we get from the top of the minIndexes stack
//
// Space O(n), we involve additional stack for keeping history of min values.
type MinStack struct {
	stack,
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (s *MinStack) Push(number int) {
	s.stack = append(s.stack, number)

	if len(s.minStack) == 0 {
		s.minStack = append(s.minStack, number)

		return
	}

	s.minStack = append(s.minStack, min(s.minStack[len(s.minStack)-1], number))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
