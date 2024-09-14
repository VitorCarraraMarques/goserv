package subway

type pair struct {
	n Node
	p []Node
}

func FindPaths(source int, dest int) [][]Node {
    var res [][]Node

    n_map := MapAllNodes()
	n_adj := MakeAdjacencyList()

	visited := make(map[int]bool)
	var queue []pair
	queue = append(queue, pair{n_map[source], []Node{n_map[source]}})
	visited[source] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		neighbors := n_adj[cur.n.Id]
		for _, neighbor := range neighbors {
			path := make([]Node, len(cur.p))
			copy(path, cur.p)
			path = append(path, neighbor)

			if neighbor.Id == dest {
		        res = append(res, path)
			} else if visited[neighbor.Id] == false {
                queue = append(queue, pair{n: neighbor, p: path})
            }

			visited[neighbor.Id] = true
		}
	}

	return res
}
