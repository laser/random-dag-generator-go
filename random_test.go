package dag_test

import (
	"math/rand"
	"testing"
	"time"

	dag "github.com/laser/random-dag-generator-go"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestPartition(t *testing.T) {
	t.Run("when nodes are connected, it returns one partition (Scenario A)", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			edges := []dag.Edge{
				{"4", "3"},
				{"3", "1"},
				{"1", "0"},
				{"3", "0"},
				{"3", "2"},
				{"0", "2"},
			}

			partitions := dag.Partition(edges)
			if len(partitions) != 1 {
				t.Errorf("Expected 1 partition, got %d", len(partitions))
			}
		}
	})

	t.Run("when nodes are connected, it returns one partition (Scenario B)", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			edges := []dag.Edge{
				{"4", "3"},
				{"3", "1"},
				{"1", "0"},
				{"3", "0"},
				{"3", "2"},
				{"0", "2"},
			}

			partitions := dag.Partition(edges)
			if len(partitions) != 1 {
				t.Errorf("Expected 1 partition, got %d", len(partitions))
			}
		}
	})

	t.Run("when two distinct graphs are present, it returns two partitions", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			edges := []dag.Edge{
				{"4", "3"},
				{"3", "1"},
				{"1", "0"},
				{"3", "0"},
				{"3", "2"},
				{"10", "20"},
				{"20", "21"},
			}

			partitions := dag.Partition(edges)
			if len(partitions) != 2 {
				t.Errorf("Expected 2 partition, got %d", len(partitions))
			}
		}
	})
}

func TestRandom(t *testing.T) {
	t.Run("it never generates a graph containing cycles", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			nodeQty := 1 + rng.Intn(100)
			maxOutdegree := 1 + rng.Intn(10)
			edgeFactor := 1.0 - rng.Float64()

			graph := dag.Random(dag.WithNodeQty(nodeQty), dag.WithMaxOutdegree(maxOutdegree), dag.WithEdgeFactor(edgeFactor))

			if containsCycle(graph) {
				t.Errorf("expected graph to be acyclic, but it contains cycles: %s", graph)
			}
		}
	})

	t.Run("it never generates an isolated node", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			nodeQty := 1 + rng.Intn(100)
			maxOutdegree := 1 + rng.Intn(10)
			edgeFactor := 1.0 - rng.Float64()

			graph := dag.Random(dag.WithNodeQty(nodeQty), dag.WithMaxOutdegree(maxOutdegree), dag.WithEdgeFactor(edgeFactor))

			a := make(map[dag.NodeId]bool)
			b := make(map[dag.NodeId]bool)

			for _, edge := range graph.Edges {
				a[edge.SourceNodeId] = true
				b[edge.TargetNodeId] = true
			}

			if len(graph.Nodes) > 1 {
				for _, node := range graph.Nodes {
					found := a[node.Id] || b[node.Id]
					if !found {
						t.Errorf("expected node %s to be connected to at least one other node", node)
					}
				}
			}
		}
	})

	t.Run("it never generates more than one graph", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			nodeQty := 1 + rng.Intn(100)
			maxOutdegree := 1 + rng.Intn(10)
			edgeFactor := 1.0 - rng.Float64()

			graph := dag.Random(dag.WithNodeQty(nodeQty), dag.WithMaxOutdegree(maxOutdegree), dag.WithEdgeFactor(edgeFactor))

			if len(graph.Nodes) > 1 && len(dag.Partition(graph.Edges)) != 1 {
				t.Errorf("expected 1 graph, got %d", len(dag.Partition(graph.Edges)))
			}
		}
	})

	t.Run("it never generates a node that violates the maximum outdegree", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			nodeQty := 1 + rng.Intn(100)
			maxOutdegree := 1 + rng.Intn(10)
			edgeFactor := 1.0 - rng.Float64()

			graph := dag.Random(dag.WithNodeQty(nodeQty), dag.WithMaxOutdegree(maxOutdegree), dag.WithEdgeFactor(edgeFactor))

			outdegrees := make(map[dag.NodeId]int)

			for _, edge := range graph.Edges {
				outdegrees[edge.SourceNodeId]++
				if outdegrees[edge.SourceNodeId] > maxOutdegree {
					t.Errorf("expected node %s to have outdegree <= %d, but it has %d", edge.SourceNodeId, maxOutdegree, outdegrees[edge.SourceNodeId])
				}
			}
		}
	})
}

func containsCycle(graph dag.Graph) bool {
	// build adjacency list
	children := make(map[dag.NodeId][]dag.NodeId)
	for _, edge := range graph.Edges {
		children[edge.SourceNodeId] = append(children[edge.SourceNodeId], edge.TargetNodeId)
	}

	// keep track of visited and 'recursion' nodes
	visited := make(map[dag.NodeId]bool)
	inStack := make(map[dag.NodeId]bool)

	// perform depth-first search to detect cycle
	for _, node := range graph.Nodes {
		if dfs(node.Id, children, visited, inStack) {
			return true
		}
	}

	return false
}

func dfs(id dag.NodeId, children map[dag.NodeId][]dag.NodeId, visited map[dag.NodeId]bool, inStack map[dag.NodeId]bool) bool {
	if inStack[id] {
		return true
	}

	if visited[id] {
		return false
	}

	visited[id] = true
	inStack[id] = true

	for _, child := range children[id] {
		if dfs(child, children, visited, inStack) {
			return true
		}
	}

	inStack[id] = false

	return false
}
