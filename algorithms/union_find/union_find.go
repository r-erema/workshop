package union_find

func UnionFind(edges []int) int {
	parents, rank := make([]int, len(edges)), make([]int, len(edges)) // parents[child] = parent

	for i := range parents {
		parents[i] = i
	}

	var find func(int) int

	find = func(vertice int) int {
		if parents[vertice] == vertice {
			return vertice
		}

		parents[vertice] = find(parents[vertice])

		return parents[vertice]
	}

	cycles := 0

	union := func(parent, child int) {
		parent1, parent2 := find(parent), find(child)

		if parent1 != parent2 {
			switch {
			case rank[parent1] < rank[parent2]:
				parents[parent1] = parent2
			case rank[parent1] > rank[parent2]:
				parents[parent2] = parent1
			default:
				parents[parent1] = parent2
				rank[parent2]++
			}

			return
		}

		cycles++
	}

	for child := range edges {
		parent := edges[child]

		union(parent, child)
	}

	return cycles
}
