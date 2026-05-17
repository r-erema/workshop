package kthlargestelementinastream_test

import (
	"testing"

	kthlargestelementinastream "github.com/r-erema/workshop/algorithms/kth_largest_element_in_a_stream"
	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	t.Parallel()

	obj := kthlargestelementinastream.Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, obj.Add(3))
	assert.Equal(t, 5, obj.Add(5))
	assert.Equal(t, 5, obj.Add(10))
	assert.Equal(t, 8, obj.Add(9))
	assert.Equal(t, 8, obj.Add(4))
}
