package main

import (
	"fmt"
	"math"
)

// Edge represents an edge in the graph
type Edge struct {
    to, weight int
}

// Graph represents a graph as an adjacency list
type Graph []*[]Edge

// BFS finds the shortest path from the source vertex to the destination vertex
// in the graph using breadth-first search
func BFS(graph Graph, source, destination int) int {
    // Initialize a queue and a distance array
    queue := []int{source}
    distances := make([]int, len(graph))
    for i := range distances {
        distances[i] = math.MaxInt32
    }
    distances[source] = 0

    // Repeat until the queue is empty
    for len(queue) > 0 {
        // Remove the first vertex from the queue
        u := queue[0]
        queue = queue[1:]

        // If the destination vertex is found, return the distance
        if u == destination {
            return distances[u]
        }

        // Add the neighbors of u to the queue
        for _, edge := range *graph[u] {
            v := edge.to
            weight := edge.weight
            if distances[v] == math.MaxInt32 {
                distances[v] = distances[u] + weight
                queue = append(queue, v)
            }
        }
    }

    return -1
}

func main() {
    // Create a graph with 5 vertices and 8 edges
    graph := make(Graph, 5)
    graph[0] = &[]Edge{
        {1, 10},
        {4, 5},
    }
    graph[1] = &[]Edge{
        {0, 10},
        {2, 1},
        {3, 2},
        {4, 5},
    }
    graph[2] = &[]Edge{
        {1, 1},
        {3, 4},
    }
    graph[3] = &[]Edge{
        {1, 2},
        {2, 4},
        {4, 2},
    }
    graph[4] = &[]Edge{
        {0, 5},
        {1, 5},
        {3, 2},
    }

    // Find the shortest path from vertex 0 to vertex 3
    distance := BFS(graph, 0, 3)
    fmt.Println(distance)  // Output: 12
}
