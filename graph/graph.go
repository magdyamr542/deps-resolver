package graph

import "errors"

type Node struct {
	Value string
}

type Edge struct {
	Src Node
	Dst Node
}

type Graph struct {
	nodes map[Node]struct{}
	edges map[Node][]Node
}

func New() Graph {
	return Graph{
		nodes: make(map[Node]struct{}),
		edges: make(map[Node][]Node),
	}
}

func (g Graph) AddNode(node Node) {
	if _, ok := g.nodes[node]; ok {
		return
	}
	g.nodes[node] = struct{}{}
}

func (g Graph) Nodes() []Node {
	nodes := []Node{}
	for node := range g.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g Graph) Neighbors(node Node) []Node {
	neighbors, ok := g.edges[node]
	if !ok {
		return nil
	}
	return neighbors
}

func (g Graph) AddEdge(src Node, dst Node) {
	if _, ok := g.nodes[src]; !ok {
		return
	}

	if _, ok := g.nodes[dst]; !ok {
		return
	}

	if g.edges[src] == nil {
		g.edges[src] = []Node{}
	}

	g.edges[src] = append(g.edges[src], dst)
}

func (g Graph) TopologicalSorting() ([]Node, error) {
	visiting := make(map[Node]struct{})
	visited := make(map[Node]struct{})
	nodes := g.Nodes()
	sorted := []Node{}

	var visitor func(Node) error
	visitor = func(node Node) error {
		// already visited. skip
		if _, ok := visited[node]; ok {
			return nil
		}

		// back to a node that we are visiting. circular dependency
		if _, ok := visiting[node]; ok {
			return errors.New("circular dependency")
		}

		// start visiting the node
		visiting[node] = struct{}{}

		// visit all neighbors of the node first. then append it to the sorted list
		neighbors := g.Neighbors(node)
		for _, neighbor := range neighbors {
			if err := visitor(neighbor); err != nil {
				return err
			}
		}

		sorted = append(sorted, node)

		// end visiting the node
		delete(visiting, node)
		visited[node] = struct{}{}

		return nil
	}

	for _, node := range nodes {
		if err := visitor(node); err != nil {
			return nil, err
		}
	}
	return sorted, nil
}
