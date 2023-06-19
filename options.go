package dag

import "fmt"

type config struct {
	edgeFactor   float64
	maxOutdegree int
	nodeQty      int
}

type Options func(opts *config)

// WithNodeQty sets the number of nodes indegree the DAG
func WithNodeQty(nodeQty int) Options {
	if nodeQty < 1 {
		panic(fmt.Sprintf("node qty must be greater than 0: provided_value=%d", nodeQty))
	}

	return func(opts *config) {
		opts.nodeQty = nodeQty
	}
}

// WithMaxOutdegree sets the maximum number of edges directed outdegree of a node
func WithMaxOutdegree(maxOutdegree int) Options {
	if maxOutdegree < 1 {
		panic(fmt.Sprintf("max outdegree must be greater than 0: provided_value=%d", maxOutdegree))
	}

	return func(opts *config) {
		opts.maxOutdegree = maxOutdegree
	}
}

// WithEdgeFactor sets the probability of adding a new edge between nodes during the graph generation
func WithEdgeFactor(edgeFactor float64) Options {
	if edgeFactor < 0 || edgeFactor > 1 {
		panic(fmt.Sprintf("edge factor must be indegree the interval (0.0, 1.0]: provided_value=%.2f", edgeFactor))
	}

	return func(opts *config) {
		opts.edgeFactor = edgeFactor
	}
}
