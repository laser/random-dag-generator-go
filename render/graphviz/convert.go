package graphviz

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
)

// RenderableGraph interface that should be implemented by different types representing graphs.
type RenderableGraph interface {
	GetNodes() []RenderableGraphNode
	GetEdges() []RenderableGraphEdge
}

// RenderableGraphNode is an interface for nodes within a Graph.
type RenderableGraphNode interface {
	GetID() string
}

// RenderableGraphEdge is an interface for edges within a Graph.
type RenderableGraphEdge interface {
	GetSourceNodeID() string
	GetTargetNodeID() string
}

// From converts a DAG to a graphviz graph
func From(dag RenderableGraph) gographviz.Graph {
	parsed, err := gographviz.ParseString("digraph {}")
	if err != nil {
		panic(err)
	}

	gvgraph := gographviz.NewGraph()

	err = gographviz.Analyse(parsed, gvgraph)
	if err != nil {
		panic(err)
	}

	// add nodes to graphviz graph
	for _, node := range dag.GetNodes() {
		err = gvgraph.AddNode("", fmt.Sprintf(`"%s"`, node.GetID()), nil)
		if err != nil {
			panic(err)
		}
	}

	// add edges, too
	for _, edge := range dag.GetEdges() {
		err = gvgraph.AddEdge(fmt.Sprintf(`"%s"`, edge.GetSourceNodeID()), fmt.Sprintf(`"%s"`, edge.GetTargetNodeID()), true, nil)
		if err != nil {
			panic(err)
		}
	}

	return *gvgraph
}
