package test

import (
	"fmt"
	"go-code-challenge/bfs"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	// Unit test for BFS function
	graph := bfs.NewGraph()
	graph.AddEdge("Alice", "Bob")
	graph.AddEdge("Alice", "Charlie")
	graph.AddEdge("Bob", "David")
	graph.AddEdge("Bob", "Eve")
	graph.AddEdge("Charlie", "Frank")
	graph.AddEdge("Charlie", "Grace")
	result := graph.BFSFind("Alice", "Grace")

	if result {
		fmt.Printf("%s ditemukan dalam graf.\n", "Grace")

	} else {
		fmt.Printf("%s tidak ditemukan dalam graf.\n", "Grace")

	}

}
