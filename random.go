package dag

import (
	"fmt"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Random generates a random DAG using the provided options
func Random(options ...Options) (out Graph) {
	cfg := config{
		nodeQty:      10,
		edgeFactor:   1.0,
		maxOutdegree: 2,
	}

	for _, option := range options {
		option(&cfg)
	}

	out.Nodes = make([]Node, cfg.nodeQty)

	for i := 0; i < cfg.nodeQty; i++ {
		out.Nodes[i] = Node{Id: NodeId(fmt.Sprintf("%d", i))}
	}

	// initialize track indegree/outdegree of each node
	degrees := make(map[NodeId]*degree)
	for _, node := range out.Nodes {
		degrees[node.Id] = &degree{}
	}

	ranks := rng.Perm(cfg.nodeQty)

	i := 0
	for i < cfg.nodeQty-1 {
		targetIndex := i + 1 + rng.Intn(cfg.nodeQty-i-1)

		ranks[targetIndex], ranks[i+1] = ranks[i+1], ranks[targetIndex]

		source := NodeId(fmt.Sprintf("%d", ranks[i]))
		target := NodeId(fmt.Sprintf("%d", ranks[targetIndex]))

		if degrees[source].outdegree < cfg.maxOutdegree {
			out.Edges = append(out.Edges, Edge{
				SourceNodeId: source,
				TargetNodeId: target,
			})

			degrees[source].incOut()
			degrees[target].incIn()
		}

		if rng.Intn(100) < 50 { // 50 is a placeholder percentage
			i += rng.Intn(cfg.nodeQty-i) + 1
		} else {
			i++
		}
	}

	edgeMap := make(map[GroupKey]bool)
	for _, edge := range out.Edges {
		edgeMap[toKey(edge.SourceNodeId, edge.TargetNodeId)] = true
	}

	for i := 0; i < cfg.nodeQty; i++ {
		for j := i + 1; j < cfg.nodeQty; j++ {
			m := cfg.edgeFactor
			n := rng.Float64()
			winner := m > n
			source := NodeId(fmt.Sprintf("%d", ranks[i]))
			target := NodeId(fmt.Sprintf("%d", ranks[j]))
			exists := edgeMap[toKey(source, target)]

			if !exists && winner && degrees[source].outdegree < cfg.maxOutdegree {
				out.Edges = append(out.Edges, Edge{
					SourceNodeId: source,
					TargetNodeId: target,
				})

				edgeMap[toKey(source, target)] = true

				degrees[source].incOut()
				degrees[target].incIn()
			}
		}
	}

	if len(out.Nodes) < 2 {
		return
	}

	// add isolated nodes to the graph
	for _, node := range out.Nodes {
		if degrees[node.Id].indegree == 0 && degrees[node.Id].outdegree == 0 {
			source := node.Id
			target := node.Id
			for target == source {
				target = out.Nodes[rng.Intn(len(out.Nodes))].Id
			}

			out.Edges = append(out.Edges, Edge{
				SourceNodeId: source,
				TargetNodeId: target,
			})

			degrees[source].incOut()
			degrees[target].incIn()
		}
	}

	// link disjoint graphs, respecting maximum outdegree
	partitions := Partition(out.Edges)
	if len(partitions) > 1 {
		for i := 1; i < len(partitions); i++ {
			for _, edge := range partitions[i] {
				// don't add edge if that would violate the maximum outdegree
				if degrees[edge.SourceNodeId].outdegree >= cfg.maxOutdegree {
					continue
				}

				// link the first valid edge in the partition to the first node in the first partition
				out.Edges = append(out.Edges, Edge{
					SourceNodeId: edge.SourceNodeId,
					TargetNodeId: partitions[0][0].SourceNodeId,
				})

				break
			}
		}
	}

	return
}

// Partition partitions a set of edges into a slice of subgraphs
func Partition(edges []Edge) (partitions [][]Edge) {
	neighbors := map[NodeId]map[NodeId]struct{}{}
	for _, e := range edges {
		setNested(neighbors, e.SourceNodeId, e.TargetNodeId, struct{}{})
		setNested(neighbors, e.TargetNodeId, e.SourceNodeId, struct{}{})
	}

	checked := map[NodeId]bool{}

	for n := range neighbors {
		if !checked[n] {
			partition := bfs(n, neighbors, checked)
			if len(partition) > 0 {
				partitions = append(partitions, partition)
			}
		}
	}

	return
}
