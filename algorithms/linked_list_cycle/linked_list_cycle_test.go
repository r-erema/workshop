package linkedlistcycle_test

import (
	"testing"

	linkedlistcycle "github.com/r-erema/workshop/algorithms/linked_list_cycle"
	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *linkedlistcycle.ListNode
		pos   int
		want  bool
	}{
		{
			name: "Simple list",
			input: func() *linkedlistcycle.ListNode {
				node1 := &linkedlistcycle.ListNode{
					Val:  3,
					Next: nil,
				}
				node2 := &linkedlistcycle.ListNode{
					Val:  2,
					Next: nil,
				}
				node3 := &linkedlistcycle.ListNode{
					Val:  0,
					Next: nil,
				}
				node4 := &linkedlistcycle.ListNode{
					Val:  -4,
					Next: nil,
				}

				node1.Next = node2
				node2.Next = node3
				node3.Next = node4
				node4.Next = node2

				return node1
			}(),
			pos:  1,
			want: true,
		},
		{
			name:  "Empty list",
			input: nil,
			pos:   -1,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, linkedlistcycle.HasCycle(tt.input))
		})
	}
}
