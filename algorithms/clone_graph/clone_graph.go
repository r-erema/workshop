package clone_graph

import "github.com/r-erema/workshop/utils/data_structure/graph"

// Time O(N+E), since we go to every Node(N) and Edge(E)
// Space O(N), since we don't store every cloned Node.

func CloneGraph(node *graph.Node) *graph.Node {
	if node == nil {
		return nil
	}

	var curr *graph.Node

	created := make(map[int]*graph.Node)
	visited := make(map[int]struct{})

	firstCloned := &graph.Node{Val: node.Val}
	res := firstCloned
	queue := []*graph.Node{node}
	created[firstCloned.Val] = firstCloned

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		if _, ok := visited[curr.Val]; ok {
			continue
		}

		visited[curr.Val] = struct{}{}

		for i := range curr.Neighbors {
			if _, ok := created[curr.Neighbors[i].Val]; !ok {
				neigh := &graph.Node{Val: curr.Neighbors[i].Val}
				created[neigh.Val] = neigh
			}

			created[curr.Val].Neighbors = append(created[curr.Val].Neighbors, created[curr.Neighbors[i].Val])

			queue = append(queue, curr.Neighbors[i])
		}
	}

	return res
}
