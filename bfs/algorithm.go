package bfs

import (
	"container/list"
	"strings"
)

func BreadthFirstSearch(tree map[string]Node, root, target string) string {
	const notFound = "not_found" // default value

	// check if root and target exist in the tree
	rootNode, rootExists := tree[root]
	_, targetExists := tree[target]
	if !rootExists || !targetExists {
		return notFound
	}

	// initialize the queue and push the root node
	q := list.New()
	q.PushBack(rootNode)

	// create a parent map to save the interactions and recreate the path
	parents := make(map[string]string) // initialize queue
	parents[root] = ""                 // initialize root without any parents

	// while queue has elements, keep iterating
	for q.Len() > 0 {
		currentNode := q.Front().Value.(Node) // get first element
		q.Remove(q.Front())                   // remove first element from queue

		// compare if node is equals to target
		if strings.EqualFold(currentNode.Value, target) {
			// the target has been looked
			// reconstructing the path
			var route []string
			for len(currentNode.Value) > 0 {
				// recreating route
				route = append([]string{currentNode.Value}, route...)
				// changing pointer
				currentNode.Value = parents[currentNode.Value]
			}

			// returning path result
			return strings.Join(route, "->")
		}

		// iterate neighbors
		for _, neighbor := range currentNode.Neighbors {
			// check if the neighbor has not already been visited
			if _, visited := parents[neighbor]; !visited {
				parents[neighbor] = currentNode.Value // add neighbor to parents map associated to current node value
				q.PushBack(tree[neighbor])            // add neighbor to the end of the queue
			}
		}
	}

	return notFound
}
