package dag

import (
	"fmt"
	"strings"
)

const SEPARATOR = "->"

type GroupKey string

func bfs(start NodeId, neighbors map[NodeId]map[NodeId]struct{}, visited map[NodeId]bool) []Edge {
	group := map[GroupKey]bool{}
	stack := []NodeId{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[node] {
			visited[node] = true
			for n := range neighbors[node] {
				group[toKey(node, n)] = true
				stack = append(stack, n)
			}
		}
	}

	edges := []Edge{}
	for k := range group {
		s, t := fromKey(k)
		edges = append(edges, Edge{SourceNodeId: s, TargetNodeId: t})
	}

	return edges
}

func toKey(source, target NodeId) GroupKey {
	return GroupKey(fmt.Sprintf("%s%s%s", source, SEPARATOR, target))
}

func fromKey(key GroupKey) (source, target NodeId) {
	parts := strings.Split(string(key), SEPARATOR)
	return NodeId(parts[0]), NodeId(parts[1])
}

func setNested[T comparable, U comparable, V any](m map[T]map[U]V, x T, y U, z V) {
	if _, ok := m[x]; !ok {
		m[x] = make(map[U]V)
	}
	m[x][y] = z
}

type degree struct {
	indegree  int
	outdegree int
}

func (d *degree) incIn() {
	d.indegree++
}

func (d *degree) incOut() {
	d.outdegree++
}
