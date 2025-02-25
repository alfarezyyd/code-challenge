package bidirectional

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Struktur Graph
type Graph struct {
	AdjacencyList map[string][]string
}

// Menambahkan edge ke graph
func (graphInstance *Graph) AddEdge(u, v string) {
	graphInstance.AdjacencyList[u] = append(graphInstance.AdjacencyList[u], v)
	graphInstance.AdjacencyList[v] = append(graphInstance.AdjacencyList[v], u) // Bidirectional graph
}

// BFS dengan Go Concurrency
func BidirectionalSearch(graphInstance *Graph, start, goal string) []string {
	if start == goal {
		return []string{start}
	}

	// Queue dan visited untuk dua arah pencarian
	startQueue := make(chan string, 10000)
	goalQueue := make(chan string, 10000)
	startVisited := make(map[string]string) // Key = node, Value = parent
	goalVisited := make(map[string]string)

	// Mutex dan WaitGroup
	var mu sync.Mutex
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) // Context untuk menghentikan goroutine jika ditemukan

	// Menandai node awal
	startQueue <- start
	startVisited[start] = ""
	goalQueue <- goal
	goalVisited[goal] = ""

	found := make(chan string, 1) // Channel untuk hasil jika ditemukan

	// BFS dengan concurrency
	bfs := func(queue chan string, visited map[string]string, otherVisited map[string]string, name string) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case node, ok := <-queue:
				if !ok {
					return
				}
				time.Sleep(1 * time.Second)

				fmt.Println("BFS", name, "memproses:", node) // Debugging

				mu.Lock()
				if _, ok := otherVisited[node]; ok {
					fmt.Println("BFS", name, "bertemu di:", node) // Debugging
					found <- node
					cancel()
					mu.Unlock()
					return
				}
				mu.Unlock()

				// Jelajahi neighbor
				for _, neighbor := range graphInstance.AdjacencyList[node] {
					mu.Lock()
					if _, seen := visited[neighbor]; !seen {
						visited[neighbor] = node
						select {
						case queue <- neighbor:
						default:
						}
					}
					mu.Unlock()
				}
			}
		}
	}

	// Jalankan BFS dengan nama
	wg.Add(2)
	go bfs(startQueue, startVisited, goalVisited, "START")
	go bfs(goalQueue, goalVisited, startVisited, "GOAL")
	// Jalankan BFS secara paralel

	// Tunggu sampai ada jalur yang ditemukan atau tidak ada jalur
	go func() {
		wg.Wait()
		close(found)
	}()

	// Rekonstruksi jalur jika ditemukan
	midpoint, ok := <-found
	fmt.Println(midpoint, ok)
	if !ok {
		return nil // Tidak ada jalur
	}

	// Bangun jalur dari start ke midpoint
	var path []string
	for node := midpoint; node != ""; node = startVisited[node] {
		path = append([]string{node}, path...)
	}

	// Bangun jalur dari midpoint ke goal
	for node := goalVisited[midpoint]; node != ""; node = goalVisited[node] {
		path = append(path, node)
	}

	return path
}
