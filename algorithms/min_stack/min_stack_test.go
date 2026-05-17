package minstack_test

import (
	"testing"

	minstack "github.com/r-erema/workshop/algorithms/min_stack"
	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	t.Parallel()

	stack := minstack.Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Equal(t, -3, stack.GetMin())
	stack.Pop()
	assert.Equal(t, 0, stack.Top())
	assert.Equal(t, -2, stack.GetMin())
}
