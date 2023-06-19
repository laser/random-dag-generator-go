package dag

type Node struct {
	Id NodeId
}

type NodeId string

type Edge struct {
	SourceNodeId NodeId
	TargetNodeId NodeId
}

type Graph struct {
	Nodes []Node
	Edges []Edge
}
