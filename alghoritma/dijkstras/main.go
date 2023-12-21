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

func (g *Graph) addEdge(src, dst, val int) {
	g.matrix[src][dst] = val
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

type data struct {
	distance int
	prevNode string
}

func (g *Graph) ShortestPath(src, dst int) int {
	fromNode := g.nodes[src]
	toNode := g.nodes[dst]

	visited := make(map[string]bool)
	hashMap := make(map[string]data)

	hashMap[g.nodes[src]] = data{
		distance: 0,
		prevNode: "",
	}

	for len(visited) < len(g.nodes) {
		visited[g.nodes[src]] = true
		fmt.Printf("%s visited \n", g.nodes[src])

		minDist := -1
		minNode := 0
		for i := 0; i < len(g.matrix[src]); i++ {
			if g.matrix[src][i] > 0 {

				d, ok := hashMap[g.nodes[i]]
				if !ok || d.distance > (g.matrix[src][i]+hashMap[g.nodes[src]].distance) {
					hashMap[g.nodes[i]] = data{
						distance: g.matrix[src][i] + hashMap[g.nodes[src]].distance,
						prevNode: g.nodes[src],
					}
				}

				if !visited[g.nodes[i]] && (minDist > hashMap[g.nodes[i]].distance || minDist == -1) {
					minDist = hashMap[g.nodes[i]].distance
					minNode = i
				}

			}
		}
		if g.nodes[src] == toNode {
			for i := 0; i < len(g.nodes); i++ {
				if !visited[g.nodes[i]] {
					src = i
				}
			}
		} else {
			src = minNode
		}
	}

	nodes := make([]string, 0, len(g.nodes))
	nodes = append(nodes, toNode)
	for fromNode != toNode {
		d := hashMap[toNode]
		toNode = d.prevNode
		nodes = append(nodes, toNode)
	}

	fmt.Println("Path:", nodes)
	return hashMap[g.nodes[dst]].distance
}

func main() {
	graph := NewGraph(6)
	graph.addNode("A")
	graph.addNode("B")
	graph.addNode("C")
	graph.addNode("D")
	graph.addNode("E")
	graph.addNode("F")

	graph.addEdge(0, 1, 2)
	graph.addEdge(0, 2, 8)
	graph.addEdge(1, 2, 5)
	graph.addEdge(1, 3, 6)
	graph.addEdge(2, 1, 5)
	graph.addEdge(2, 3, 3)
	graph.addEdge(2, 4, 2)
	graph.addEdge(3, 5, 9)
	graph.addEdge(4, 5, 3)

	graph.print()

	s := graph.ShortestPath(0, 5)
	fmt.Println(s)
}
