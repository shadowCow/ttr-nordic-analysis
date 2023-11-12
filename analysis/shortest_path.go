package analysis

import (
	"math"
)

type Graph struct {
	// city to city to one-hop cost
	Edges    map[City]map[City]int
}

func NewGraph(routes []Route) *Graph {
	graph := &Graph{
		Edges:    make(map[City]map[City]int),
	}

	for _, route := range routes {
		// some city pairs have multiple routes
		// just keep the shorter of the two
		f, fromExists := graph.Edges[route.From]
		if fromExists {
			cost, toExists := f[route.To]
			if toExists {
				if route.Length > cost {
					continue
				}
			}
		}

		graph.addEdge(route.From, route.To, route.Length)
		graph.addEdge(route.To, route.From, route.Length)
	}

	return graph
}

func (g *Graph) addEdge(src, dest City, weight int) {
	if g.Edges[src] == nil {
		g.Edges[src] = make(map[City]int)
	}
	g.Edges[src][dest] = weight
}

// returns the cost of the shortest path
// and the path through all intermediate cities.
func (g *Graph) ShortestPath(from, to City) (int, []City) {
	shortestPathCosts, shortestPaths := dijkstra(g, from)

	return shortestPathCosts[to], shortestPaths[to]
}

// returns...
// 1 - a map that has the shortest cost of getting to every City
//     from the start City.
//     If the start city is 'Oslo', then distances['Stockholm']
//     will give the shortest cost from Oslo to Stockholm
// 2 - a map with the path taken for the shortest path.
func dijkstra(g *Graph, start City) (map[City]int, map[City][]City) {
	distances := make(map[City]int)
	subpaths := make(map[City][]City)
	visited := make(map[City]bool)

	// Initialize distances and subpaths
	for vertex := range g.Edges {
		distances[vertex] = math.MaxInt32
		subpaths[vertex] = []City{}
	}
	distances[start] = 0

	// Priority queue to keep track of the nodes with the shortest distance
	queue := make([]City, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visited[current] = true

		for neighbor, weight := range g.Edges[current] {
			// Calculate the distance from the start node to the neighbor node
			newDistance := distances[current] + weight

			// If a shorter path is found, update distances and subpaths
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				subpaths[neighbor] = append(subpaths[current], neighbor)

				// only push unvisited nodes onto the queue
				if _, exists := visited[neighbor]; !exists {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return distances, subpaths
}
