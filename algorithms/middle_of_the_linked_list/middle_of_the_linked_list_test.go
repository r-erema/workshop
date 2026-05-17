package middleofthelinkedlist_test

import (
	"testing"

	middleofthelinkedlist "github.com/r-erema/workshop/algorithms/middle_of_the_linked_list"
	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input, want *middleofthelinkedlist.ListNode
	}{
		{
			name: "5 nodes list",
			input: &middleofthelinkedlist.ListNode{
				Val: 1,
				Next: &middleofthelinkedlist.ListNode{
					Val: 2,
					Next: &middleofthelinkedlist.ListNode{
						Val: 3,
						Next: &middleofthelinkedlist.ListNode{
							Val: 4,
							Next: &middleofthelinkedlist.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			want: &middleofthelinkedlist.ListNode{
				Val: 3,
				Next: &middleofthelinkedlist.ListNode{
					Val: 4,
					Next: &middleofthelinkedlist.ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, middleofthelinkedlist.MiddleNode(tt.input))
		})
	}
}
