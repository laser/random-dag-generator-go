package graphviz

import (
	"os"
	"os/exec"

	"github.com/awalterschulze/gographviz"
)

// RenderTo renders a graphviz graph as a PNG file at the given path using the `dot` command
func RenderTo(graph gographviz.Graph, path string) {
	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}

	defer func() {
		os.Remove(file.Name())
	}()

	// write DOT file
	err = os.WriteFile(file.Name(), []byte(graph.String()), 0644)
	if err != nil {
		panic(err)
	}

	// render DOT file as PNG
	cmd := exec.Command("dot", "-Tpng", file.Name(), "-o", path)

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
