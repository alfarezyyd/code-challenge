package test

import (
	"fmt"
	"go-code-challenge/gbfs"
	"testing"
)

// TestGreedyBestFirstSearch menguji algoritma Greedy Best-First Search
func TestGreedyBestFirstSearch(t *testing.T) {
	// Inisialisasi graph dengan adjacency list
	graph := gbfs.Graph{
		AdjacencyList: map[string][]string{
			"Kebon Sirih":      {"Bundaran HI", "Kuningan"},
			"Bundaran HI":      {"Kebon Sirih", "BNI City", "Karet"},
			"Karet":            {"Bundaran HI", "Dukuh Atas"},
			"BNI City":         {"Bundaran HI", "Setiabudi", "Dukuh Atas"},
			"Setiabudi":        {"BNI City", "Flyover Kuningan", "Kuningan"},
			"Kuningan":         {"Setiabudi", "Kebon Sirih"},
			"Flyover Kuningan": {"Setiabudi", "Pasar Rumput"},
			"Pasar Rumput":     {"Flyover Kuningan", "Manggarai", "Dukuh Atas"},
			"Dukuh Atas":       {"BNI City", "Karet", "Pasar Rumput"},
			"Manggarai":        {"Pasar Rumput", "Flyover Pramuka", "Klender"},
			"Flyover Pramuka":  {"Manggarai", "Matraman", "Pasar Genjing"},
			"Matraman":         {"Flyover Pramuka", "Gunung Sahari", "Matraman Baru"},
			"Gunung Sahari":    {"Matraman", "Ancol"},
			"Ancol":            {"Gunung Sahari"},
			"Matraman Baru":    {"Matraman", "Jatinegara"},
			"Jatinegara":       {"Matraman Baru"},
			"Pasar Genjing":    {"Flyover Pramuka", "Velodrome", "Pemuda Merdeka", "Klender"},
			"Velodrome":        {"Pasar Genjing"},
			"Pemuda Merdeka":   {"Pasar Genjing", "Pulo Gadung"},
			"Pulo Gadung":      {"Pemuda Merdeka"},
			"Klender":          {"Manggarai", "Kebon Nanas", "Karet", "Pasar Genjing"},
			"Kebon Nanas":      {"Klender", "Cawang", "Cawang Sentral"},
			"Cawang":           {"Kebon Nanas"},
			"Cawang Sentral":   {"Kebon Nanas"},
		},
		Heuristic: map[string]int{
			// Nilai heuristic bisa diubah sesuai kebutuhan algoritma GBFS
			"Kebon Sirih":      13,
			"Bundaran HI":      14,
			"Karet":            13,
			"BNI City":         12,
			"Setiabudi":        11,
			"Flyover Kuningan": 10,
			"Pasar Rumput":     9,
			"Dukuh Atas":       9,
			"Manggarai":        8,
			"Flyover Pramuka":  7,
			"Matraman":         6,
			"Gunung Sahari":    5,
			"Ancol":            4,
			"Matraman Baru":    6,
			"Jatinegara":       5,
			"Kuningan":         10,
			"Pasar Genjing":    6,
			"Velodrome":        5,
			"Pemuda Merdeka":   4,
			"Pulo Gadung":      3,
			"Klender":          7,
			"Kebon Nanas":      6,
			"Cawang":           5,
			"Cawang Sentral":   5,
		},
	}

	// Jalankan pencarian dari A ke G
	path := graph.GreedyBestFirstSearch("Pulo Gadung", "Kebon Sirih")

	// Jalur yang diharapkan berdasarkan heuristic
	expectedPath := []string{"Pulo Gadung", "Pemuda Merdeka", "Pasar Genjing", "Klender", "Manggarai", "Pasar Rumput", "Flyover Kuningan", "Setiabudi", "Kuningan", "Kebon Sirih"}

	// Cek apakah hasilnya sesuai dengan yang diharapkan
	if len(path) != len(expectedPath) {
		t.Errorf("Expected path length %d, but got %d", len(expectedPath), len(path))
	}

	for i := range expectedPath {
		if path[i] != expectedPath[i] {
			t.Errorf("Expected path %v, but got %v", expectedPath, path)
			break
		}
	}

	fmt.Println(path, expectedPath)
}
