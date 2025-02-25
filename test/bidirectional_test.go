package test

import (
	"go-code-challenge/bidirectional"
	"testing"
)

// Test untuk Bidirectional BFS
func TestBidirectionalBFS(t *testing.T) {
	// Buat graph contoh
	graph := &bidirectional.Graph{AdjacencyList: make(map[string][]string)}
	graph.AddEdge("Manggarai", "Pasar Rumput")
	graph.AddEdge("Pasar Rumput", "Manggarai")
	// Pasar Rumput
	// Pasar Rumput - Flyover Kuningan
	graph.AddEdge("Pasar Rumput", "Flyover Kuningan")
	graph.AddEdge("Flyover Kuningan", "Setiabudi")
	graph.AddEdge("Setiabudi", "Flyover Kuningan")
	graph.AddEdge("Setiabudi", "Kuningan")

	// Pasar Rumput - Dukuh Atas
	graph.AddEdge("Pasar Rumput", "Dukuh Atas")
	graph.AddEdge("Dukuh Atas", "Karet")

	// Pasar Rumput - Dukuh Atas - BNI
	graph.AddEdge("Dukuh Atas", "BNI City")
	graph.AddEdge("BNI City", "Bundaran HI")
	graph.AddEdge("Bundaran HI", "Kebon Sirih")

	// Klender
	graph.AddEdge("Manggarai", "Klender")
	graph.AddEdge("Klender", "Kebon Nanas")
	graph.AddEdge("Kebon Nanas", "Cawang")
	graph.AddEdge("Kebon Nanas", "Cawang Sentral")

	// Flyover Pramuka
	graph.AddEdge("Manggarai", "Flyover Pramuka")

	// Flyover Pramuka - Matraman
	graph.AddEdge("Flyover Pramuka", "Matraman")
	// Flyover Pramuka - Matraman - Gunung Sahari
	graph.AddEdge("Matraman", "Gunung Sahari")
	graph.AddEdge("Gunung Sahari", "Ancol")

	// Flyover Pramuka - Matraman - Matraman Baru
	graph.AddEdge("Matraman", "Matraman Baru")
	graph.AddEdge("Matraman Baru", "Jatinegara")

	// Flyover Pramuka - Pasar Genjing
	graph.AddEdge("Flyover Pramuka", "Pasar Genjing")

	// Flyover Pramuka - Pasar Genjing - Velodrome
	graph.AddEdge("Pasar Genjing", "Velodrome")

	// Flyover Pramuka - Pasar Genjing - Pemuda Merdeka
	graph.AddEdge("Pasar Genjing", "Pemuda Merdeka")
	graph.AddEdge("Pemuda Merdeka", "Pulo Gadung")

	// Daftar test case
	tests := []struct {
		start, goal string
		expected    []string
	}{
		{"Kebon Sirih", "Pulo Gadung", []string{"Kebon Sirih", "Bundaran HI", "BNI City", "Dukuh Atas", "Pasar Rumput", "Manggarai", "Flyover Pramuka", "Pasar Genjing", "Pemuda Merdeka", "Pulo Gadung"}}, // Jalur dari A ke G
	}

	// Loop melalui setiap test case
	for _, tt := range tests {
		t.Run(tt.start+" to "+tt.goal, func(t *testing.T) {
			result := bidirectional.BidirectionalSearch(graph, tt.start, tt.goal)
			if !equalSlice(result, tt.expected) {
				t.Errorf("Path from %s to %s: got %v, expected %v", tt.start, tt.goal, result, tt.expected)
			}
		})
	}
}

// Fungsi helper untuk membandingkan slice
func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
