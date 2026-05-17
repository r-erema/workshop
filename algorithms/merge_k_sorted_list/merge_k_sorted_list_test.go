package mergeksortedlist_test

import (
	"testing"

	mergeksortedlist "github.com/r-erema/workshop/algorithms/merge_k_sorted_list"
	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		lists []*mergeksortedlist.ListNode
		want  *mergeksortedlist.ListNode
	}{
		{
			name: "Case 0",
			lists: []*mergeksortedlist.ListNode{
				{
					Value: 1,
					Next: &mergeksortedlist.ListNode{
						Value: 4,
						Next: &mergeksortedlist.ListNode{
							Value: 5,
							Next:  nil,
						},
					},
				},
				{
					Value: 1,
					Next: &mergeksortedlist.ListNode{
						Value: 3,
						Next: &mergeksortedlist.ListNode{
							Value: 4,
							Next:  nil,
						},
					},
				},
				{
					Value: 2,
					Next: &mergeksortedlist.ListNode{
						Value: 6,
						Next:  nil,
					},
				},
			},
			want: &mergeksortedlist.ListNode{
				Value: 1,
				Next: &mergeksortedlist.ListNode{
					Value: 1,
					Next: &mergeksortedlist.ListNode{
						Value: 2,
						Next: &mergeksortedlist.ListNode{
							Value: 3,
							Next: &mergeksortedlist.ListNode{
								Value: 4,
								Next: &mergeksortedlist.ListNode{
									Value: 4,
									Next: &mergeksortedlist.ListNode{
										Value: 5,
										Next: &mergeksortedlist.ListNode{
											Value: 6,
											Next:  nil,
										},
									},
								},
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

			assert.Equal(t, tt.want, mergeksortedlist.MergeKSortedLists(tt.lists))
		})
	}
}
