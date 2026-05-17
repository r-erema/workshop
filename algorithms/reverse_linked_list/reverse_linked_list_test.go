package reverselinkedlist_test

import (
	"testing"

	reverselinkedlist "github.com/r-erema/workshop/algorithms/reverse_linked_list"
	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *reverselinkedlist.ListNode
		want  *reverselinkedlist.ListNode
	}{
		{
			name: "Simple list",
			input: &reverselinkedlist.ListNode{
				Val: 1,
				Next: &reverselinkedlist.ListNode{
					Val: 2,
					Next: &reverselinkedlist.ListNode{
						Val: 3,
						Next: &reverselinkedlist.ListNode{
							Val: 4,
							Next: &reverselinkedlist.ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: &reverselinkedlist.ListNode{
				Val: 5,
				Next: &reverselinkedlist.ListNode{
					Val: 4,
					Next: &reverselinkedlist.ListNode{
						Val: 3,
						Next: &reverselinkedlist.ListNode{
							Val: 2,
							Next: &reverselinkedlist.ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, reverselinkedlist.ReverseList(tt.input))
		})
	}
}
