package copylistwithrandompointer_test

import (
	"testing"

	copylistwithrandompointer "github.com/r-erema/workshop/algorithms/copy_list_with_random_pointer"
	linkedlist "github.com/r-erema/workshop/utils/data_structure/linked_list"
	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		buildList func() *linkedlist.Node
	}{
		{
			name: "2 nodes list",
			buildList: func() *linkedlist.Node {
				node1 := &linkedlist.Node{
					Val:  8,
					Next: nil,
					Prev: nil,
				}

				node2 := &linkedlist.Node{
					Val:  11,
					Next: nil,
					Prev: nil,
				}

				node1.Next = node2
				node1.Prev = node2

				node2.Next = nil
				node2.Prev = node1

				return node1
			},
		},
		{
			name: "5 nodes list",
			buildList: func() *linkedlist.Node {
				node1 := &linkedlist.Node{
					Val:  7,
					Next: nil,
					Prev: nil,
				}

				node2 := &linkedlist.Node{
					Val:  13,
					Next: nil,
					Prev: nil,
				}

				node3 := &linkedlist.Node{
					Val:  13,
					Next: nil,
					Prev: nil,
				}

				node4 := &linkedlist.Node{
					Val:  13,
					Next: nil,
					Prev: nil,
				}

				node5 := &linkedlist.Node{
					Val:  13,
					Next: nil,
					Prev: nil,
				}

				node1.Next = node2
				node1.Prev = nil

				node2.Next = node3
				node2.Prev = node1

				node3.Next = node4
				node3.Prev = node5

				node4.Next = node5
				node4.Prev = node3

				node5.Next = nil
				node5.Prev = node1

				return node1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			head := tt.buildList()
			copied := copylistwithrandompointer.CopyPrevList(head)
			assert.True(t, copylistwithrandompointer.IsClone(head, copied))
		})
	}
}
