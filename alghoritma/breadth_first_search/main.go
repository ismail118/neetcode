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

type Queue struct {
	data []int
}

func (q *Queue) offer(src int) {
	q.data = append(q.data, src)
}

func (q *Queue) poll() int {
	x := q.data[0]

	q.data = q.data[1:]

	return x
}

func (q *Queue) size() int {
	return len(q.data)
}

// BFS
func (g *Graph) breadthFirstSearch(src int) {

	queue := &Queue{
		data: make([]int, 0),
	}
	visited := make([]bool, len(g.matrix))

	queue.offer(src)
	visited[src] = true

	for queue.size() != 0 {
		src = queue.poll()
		fmt.Printf("%s visited \n", g.nodes[src])

		for i := 0; i < len(g.matrix[src]); i++ {
			if g.matrix[src][i] == 1 && !visited[i] {
				queue.offer(i)
				visited[i] = true
			}
		}
	}
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

	graph.breadthFirstSearch(0)

}
