package bfs

import (
	"bytes"
)

type Graph struct {
	adjacencyList map[string][]string
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
	}
}

// AddEdge adds an edge between two people
func (g *Graph) AddEdge(person1, person2 string) {
	g.adjacencyList[person1] = append(g.adjacencyList[person1], person2)
}

// BFS performs Breadth-First Search starting from a given person
func (g *Graph) BFS(start string) string {
	visited := make(map[string]bool)
	queue := []string{start}
	var result bytes.Buffer

	visited[start] = true

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]
		result.WriteString(person + " ")

		for _, neighbor := range g.adjacencyList[person] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return result.String()
}

func (g *Graph) BFSFind(start, target string) bool {
	visited := make(map[string]bool)
	queue := []string{start}

	visited[start] = true

	for len(queue) > 0 {
		person := queue[0]
		queue = queue[1:]

		// Jika menemukan orang yang dicari, langsung return true
		if person == target {
			return true
		}

		for _, neighbor := range g.adjacencyList[person] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// Jika tidak ditemukan, return false
	return false
}
