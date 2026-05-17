package addtwonumbers_test

import (
	"testing"

	addtwonumbers "github.com/r-erema/workshop/algorithms/add_two_numbers"
	linkedlist "github.com/r-erema/workshop/utils/data_structure/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		list1, list2, want *linkedlist.Node
	}{
		{
			name: "addition with moving to the next rank",
			list1: &linkedlist.Node{
				Val: 2,
				Next: &linkedlist.Node{
					Val: 4,
					Next: &linkedlist.Node{
						Val: 3,
					},
				},
			},
			list2: &linkedlist.Node{
				Val: 5,
				Next: &linkedlist.Node{
					Val: 6,
					Next: &linkedlist.Node{
						Val: 4,
					},
				},
			},
			want: &linkedlist.Node{
				Val: 7,
				Next: &linkedlist.Node{
					Val: 0,
					Next: &linkedlist.Node{
						Val: 8,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, addtwonumbers.AddTwoNumbers(tt.list1, tt.list2))
		})
	}
}
