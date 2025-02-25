package gbfs

import (
	"container/heap"
	"fmt"
)

// Struktur untuk merepresentasikan graf
type Graph struct {
	AdjacencyList map[string][]string
	Heuristic     map[string]int // Nilai heuristik tiap node
}

// Struktur untuk antrian prioritas
type PriorityQueue []Node

type Node struct {
	Name     string
	Priority int // Nilai heuristik (h(n))
	Index    int
}

// Implementasi heap.Interface untuk PriorityQueue
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(Node)
	node.Index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// Fungsi Greedy Best-First Search
func (graph *Graph) GreedyBestFirstSearch(start, goal string) []string {
	// Priority queue untuk menyimpan node yang akan dieksplorasi
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{Name: start, Priority: graph.Heuristic[start], Index: 0})

	// Menyimpan jalur yang diambil
	cameFrom := make(map[string]string)

	// Track nodes that are in the frontier (queue)
	inFrontier := make(map[string]bool)
	inFrontier[start] = true

	// Track nodes that have been visited/explored
	visited := make(map[string]bool)

	fmt.Println("Starting search from", start, "to", goal)
	fmt.Println("Initial heuristic:", graph.Heuristic[start])

	for pq.Len() > 0 {
		// Ambil node dengan nilai heuristik terendah
		current := heap.Pop(pq).(Node)

		fmt.Printf("Exploring: %s (h=%d)\n", current.Name, graph.Heuristic[current.Name])

		// Skip if we've already processed this node
		if visited[current.Name] {
			fmt.Printf("  Already visited %s, skipping\n", current.Name)
			continue
		}

		// Mark as visited
		visited[current.Name] = true
		inFrontier[current.Name] = false // No longer in frontier

		// Check if we've reached the goal
		if current.Name == goal {
			fmt.Println("Goal reached!")
			path := reconstructPath(cameFrom, start, goal)
			fmt.Println("Final path:", path)
			return path
		}

		// Process neighbors
		fmt.Printf("  Neighbors of %s:\n", current.Name)
		for _, neighbor := range graph.AdjacencyList[current.Name] {
			fmt.Printf("    Checking neighbor: %s (h=%d)\n", neighbor, graph.Heuristic[neighbor])

			if visited[neighbor] {
				fmt.Printf("      Already visited %s, skipping\n", neighbor)
				continue
			}

			if !inFrontier[neighbor] {
				// If not in frontier yet, add it
				cameFrom[neighbor] = current.Name
				heap.Push(pq, Node{Name: neighbor, Priority: graph.Heuristic[neighbor]})
				inFrontier[neighbor] = true
				fmt.Printf("      Added %s to frontier with h=%d\n", neighbor, graph.Heuristic[neighbor])
			}
		}

		// Print current state of priority queue for debugging
		fmt.Printf("  Current frontier: ")
		tempPQ := make(PriorityQueue, pq.Len())
		copy(tempPQ, *pq)
		for i := 0; i < len(tempPQ); i++ {
			n := heap.Pop(&tempPQ).(Node)
			fmt.Printf("%s(h=%d) ", n.Name, n.Priority)
		}
		fmt.Println()
	}

	fmt.Println("No path found")
	return nil // Tidak ditemukan jalur
}

// Fungsi untuk merekonstruksi jalur dari start ke goal
// Helper function to reconstruct the path
func reconstructPath(cameFrom map[string]string, start, goal string) []string {
	path := []string{goal}
	current := goal

	for current != start {
		current = cameFrom[current]
		path = append([]string{current}, path...)
	}

	return path
}
