package main

import (
	"fmt"
	"go-code-challenge/bfs"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	graph := bfs.NewGraph()
	graph.AddEdge("Alice", "Bob")
	graph.AddEdge("Alice", "Charlie")
	graph.AddEdge("Bob", "David")
	graph.AddEdge("Bob", "Eve")
	graph.AddEdge("Charlie", "Frank")
	graph.AddEdge("Charlie", "Grace")

	fmt.Println("BFS traversal starting from Alice:")
	fmt.Println(graph.BFS("Alice"))

}
