package graph

import "fmt"


func BFSSteps(g Graph, start string) []Step {
	steps := []Step{}
	visited := map[string]bool{}
	frontier := []string{start}
	visited[start] = true

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]
		step := Step{
			Current:  current,
			Frontier: copySlice(frontier),
			Visited:  keys(visited),
			Message:  "visiting " + current,
		}
		steps = append(steps, step)

		for _, edge := range g.Edges[current] {
			if !visited[edge.To] {
				visited[edge.To] = true
				frontier = append(frontier, edge.To)
			}
		}
	}
	return steps
}

func DFSSteps(g Graph, start string) []Step {
	steps := []Step{}
	visited := map[string]bool{}
	dfsHelper(g, start, visited, &steps)
	return steps
}

func dfsHelper(g Graph, node string, visited map[string]bool, steps *[]Step) {
	visited[node] = true

	*steps = append(*steps, Step{
		Current:  node,
		Visited:  keys(visited),
		Frontier: []string{},
		Message:  "visiting " + node,
	})

	for _, edge := range g.Edges[node] {
		if !visited[edge.To] {
			dfsHelper(g, edge.To, visited, steps)
		}
	}
}

func DijkstraSteps(g Graph, start string) []Step {
	steps := []Step{}
	dist := map[string]float64{}
	visited := map[string]bool{}
	prev := map[string]string{}

	for _, node := range g.Nodes {
		dist[node] = 1<<63 - 1
	}
	dist[start] = 0

	for len(visited) < len(g.Nodes) {
		current := minDist(dist, visited)
		if current == "" {
			break
		}
		visited[current] = true

		steps = append(steps, Step{
			Current:   current,
			Visited:   keys(visited),
			Frontier:  unvisited(g.Nodes, visited),
			EdgeTaken: copyMap(prev),
			Message:   "settled " + current + " — distance " + fmtDist(dist[current]),
		})

		for _, edge := range g.Edges[current] {
			newDist := dist[current] + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				prev[edge.To] = current
			}
		}
	}
	return steps
}

func copySlice(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	return c
}

func keys(m map[string]bool) []string {
	result := []string{}
	for k := range m {
		result = append(result, k)
	}
	return result
}

func copyMap(m map[string]string) map[string]string {
	c := map[string]string{}
	for k, v := range m {
		c[k] = v
	}
	return c
}

func unvisited(nodes []string, visited map[string]bool) []string {
	result := []string{}
	for _, n := range nodes {
		if !visited[n] {
			result = append(result, n)
		}
	}
	return result
}

func minDist(dist map[string]float64, visited map[string]bool) string {
	min := 1<<63 - 1
	node := ""
	for k, v := range dist {
		if !visited[k] && int(v) < min {
			min = int(v)
			node = k
		}
	}
	return node
}

func fmtDist(d float64) string {
	if d == 1<<63-1 {
		return "∞"
	}
	return fmt.Sprintf("%.1f", d)
}
