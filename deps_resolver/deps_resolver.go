package depsresolver

import (
	"errors"
	"fmt"

	"github.com/magdyamr542/dips/graph"
	"github.com/magdyamr542/dips/helpers"
)

// Resolver resolves dependencies between entities.
//
// Entities are represented as strings (e.g jobId in a job system).
// The resolver returns the entities in such an order that if entity A depends on entity B,
// B comes before A (e.g job B should be executed before A).
//
// If no such order can be obtained (e.g Circular dependency), an error is returned.
type Resolver interface {
	Resolve(entities []string, deps map[string][]string) ([]string, error)
}

// resolver resolves dependencies using topological sorting.
// It first creates a graph where entities are the nodes and a dependency between two entities or (nodes) is an edge.
// Then it applies topological sorting on the graph to get an order which respects the specified dependencies.
type resolver struct {
}

func NewResolver() resolver {
	return resolver{}
}

func (r resolver) Resolve(entities []string, deps map[string][]string) ([]string, error) {
	if len(entities) == 0 {
		return nil, errors.New("empty entities")
	}

	entitiesSet := helpers.NewSet(entities)
	// make sure all entities exist
	for entity, entityDeps := range deps {
		if !entitiesSet.Exists(entity) {
			return nil, fmt.Errorf("entity %q doesn't exist in the given entities", entity)
		}
		for _, dstEntity := range entityDeps {
			if !entitiesSet.Exists(dstEntity) {
				return nil, fmt.Errorf("entity %q doesn't exist in the given entities", dstEntity)
			}
		}
	}

	// build the graph
	g := graph.New()

	// nodes
	for _, entity := range entities {
		g.AddNode(graph.Node{Value: entity})
	}

	// edges
	for entity, entityDeps := range deps {
		src := graph.Node{Value: entity}
		for _, dstEntity := range entityDeps {
			dst := graph.Node{Value: dstEntity}
			g.AddEdge(src, dst)
		}
	}

	// run topological sorting
	sortedNodes, err := g.TopologicalSorting()
	if err != nil {
		return nil, fmt.Errorf("can't create dependencies: %v", err)
	}

	result := []string{}
	for _, node := range sortedNodes {
		result = append(result, node.Value)
	}

	return result, nil
}
