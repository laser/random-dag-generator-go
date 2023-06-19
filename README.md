# Random DAG Generator in Go

`random-dag-generator-go` is a library for (you guessed it!) generating random 
Directed Acyclic Graphs (DAGs) in Go.

<img src="https://github.com/laser/random-dag-generator-go/assets/884507/bb2c5672-71b0-48e0-a7d2-7e2e5dddc258" height="300">
<img src="https://github.com/laser/random-dag-generator-go/assets/884507/dbde1b23-751b-4f35-a50d-66075a91ad21" height="300">
<img src="https://github.com/laser/random-dag-generator-go/assets/884507/f84e4b52-7cf9-4060-a349-1d2fea206cf6" height="300">
<img src="https://github.com/laser/random-dag-generator-go/assets/884507/0867d4e6-6cec-494f-a693-3b0b84e13200" height="300">
<img src="https://github.com/laser/random-dag-generator-go/assets/884507/a2fbb04f-3fce-4f90-bedb-6d4af1c13f61" height="300">

## Features

- Generate random Directed Acyclic Graphs (DAGs), controlling:
  - exact quantity of nodes in the DAG
  - maximum outdegree of each node
  - probability of adding new edges during graph generation
- Render a DAG using graphviz and DOT

## Demo

```go
package main

import (
	"flag"
	"math/rand"

	dag "github.com/laser/random-dag-generator-go"
	gvz "github.com/laser/random-dag-generator-go/render/graphviz"
)

func main() {
	flag.Parse()

	graph := dag.Random(
		dag.WithNodeQty(10),
		dag.WithMaxOutdegree(3),
		dag.WithEdgeFactor(0.5))

	converted := gvz.From(graph)

	gvz.RenderTo(converted, "/tmp/flarp.png")
}
```
