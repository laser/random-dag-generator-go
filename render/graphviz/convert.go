package graphviz

import (
	"github.com/awalterschulze/gographviz"
	dag "github.com/laser/random-dag-generator-go"
)

// From converts a DAG to a graphviz graph
func From(dag dag.Graph) gographviz.Graph {
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
	for _, node := range dag.Nodes {
		err = gvgraph.AddNode("", string(node.Id), nil)
		if err != nil {
			panic(err)
		}
	}

	// add edges, too
	for _, edge := range dag.Edges {
		err = gvgraph.AddEdge(string(edge.SourceNodeId), string(edge.TargetNodeId), true, nil)
		if err != nil {
			panic(err)
		}
	}

	return *gvgraph
}
