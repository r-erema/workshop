package reorderlist_test

import (
	"testing"

	reorderlist "github.com/r-erema/workshop/algorithms/reorder_list"
	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *reorderlist.ListNode
		want  *reorderlist.ListNode
	}{
		{
			name: "Even count of nodes",
			input: &reorderlist.ListNode{
				Val: 1,
				Next: &reorderlist.ListNode{
					Val: 2,
					Next: &reorderlist.ListNode{
						Val: 3,
						Next: &reorderlist.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &reorderlist.ListNode{
				Val: 1,
				Next: &reorderlist.ListNode{
					Val: 4,
					Next: &reorderlist.ListNode{
						Val: 2,
						Next: &reorderlist.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "Not even count of nodes",
			input: &reorderlist.ListNode{
				Val: 11,
				Next: &reorderlist.ListNode{
					Val: 8,
					Next: &reorderlist.ListNode{
						Val: 5,
						Next: &reorderlist.ListNode{
							Val: -2,
							Next: &reorderlist.ListNode{
								Val:  0,
								Next: nil,
							},
						},
					},
				},
			},
			want: &reorderlist.ListNode{
				Val: 11,
				Next: &reorderlist.ListNode{
					Val: 0,
					Next: &reorderlist.ListNode{
						Val: 8,
						Next: &reorderlist.ListNode{
							Val: -2,
							Next: &reorderlist.ListNode{
								Val:  5,
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

			reorderlist.ReorderList(tt.input)
			assert.Equal(t, tt.want, tt.input)
		})
	}
}
