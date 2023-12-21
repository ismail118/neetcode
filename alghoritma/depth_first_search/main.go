package main

import "fmt"

type Graph struct {
	nodes  []string
	matrix [][]int
}

func NewGraph(size int) Graph {
	graph := Graph{
		nodes:  make([]string, 0),
		matrix: make([][]int, size),
	}

	for i, _ := range graph.matrix {
		graph.matrix[i] = make([]int, size)
	}

	return graph
}

func (g *Graph) addNode(node string) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) addEdge(src, dst int) {
	g.matrix[src][dst] = 1
}

func (g *Graph) checkEdge(src, dst int) bool {
	if g.matrix[src][dst] == 1 {
		return true
	}

	return false
}

func (g *Graph) print() {
	fmt.Printf("\n ")
	for _, node := range g.nodes {
		fmt.Printf(" %s", node)
	}
	fmt.Println()

	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix[0]); j++ {
			if j == 0 {
				fmt.Printf("%s ", g.nodes[i])
			}
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

// DFS
func (g *Graph) depthFirstSearch(src int) {
	visited := make([]bool, len(g.matrix))
	g.dFSHelper(src, visited)

}

func (g *Graph) dFSHelper(src int, visited []bool) {
	if visited[src] {
		return
	} else {
		visited[src] = true
		fmt.Printf("%s visited \n", g.nodes[src])
	}

	for i := 0; i < len(g.matrix[src]); i++ {
		if g.matrix[src][i] == 1 {
			g.dFSHelper(i, visited)
		}
	}

	return
}

func main() {
	graph := NewGraph(5)
	graph.addNode("A")
	graph.addNode("B")
	graph.addNode("C")
	graph.addNode("D")
	graph.addNode("E")

	graph.addEdge(0, 1)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 3)
	graph.addEdge(2, 4)
	graph.addEdge(4, 0)
	graph.addEdge(4, 2)

	graph.print()

	graph.depthFirstSearch(0)
}
