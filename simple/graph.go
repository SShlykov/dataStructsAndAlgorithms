package simple_structs

import "fmt"

type Graph struct {
	vertices map[string][]string
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

func (g *Graph) AddVertex(v string) {
	if _, exists := g.vertices[v]; !exists {
		g.vertices[v] = []string{}
	}
}

func (g *Graph) AddEdge(v1, v2 string) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

func graph() {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddEdge("A", "B")

	fmt.Println(g.vertices) // map[A:[B] B:[A]]
}
