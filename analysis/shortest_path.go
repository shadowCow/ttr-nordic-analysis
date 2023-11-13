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
	queue := newSortedQueue()
	queue.add(start, 0)

	for queue.length() > 0 {
		current := queue.takeHead()
		visited[current.city] = true

		for neighbor, weight := range g.Edges[current.city] {
			// Calculate the distance from the start node to the neighbor node
			newDistance := distances[current.city] + weight

			// If a shorter path is found, update distances and subpaths
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				subpaths[neighbor] = append(subpaths[current.city], neighbor)

				// only push unvisited nodes onto the queue
				if _, exists := visited[neighbor]; !exists {
					queue.add(neighbor, distances[neighbor])
				}
			}
		}
	}

	return distances, subpaths
}

type toVisit struct {
	city City
	cost int
}
type sortedQueue struct {
	items []toVisit
}
func newSortedQueue() sortedQueue {
	return sortedQueue{
		items: []toVisit{},
	}
}
func (q *sortedQueue) add(city City, cost int) {
	if len(q.items) == 0 {
		q.items = []toVisit{
			{city, cost},
		}
	} else {
		insertionIndex := len(q.items)
		for i, item := range q.items {
			if item.cost > cost {
				insertionIndex = i
				break
			}
		}

		if insertionIndex == len(q.items) {
			q.items = append(q.items, toVisit{city, cost})
		} else {
			newItems := append([]toVisit{}, q.items[:insertionIndex]...)
				after := q.items[insertionIndex:]
				newItems = append(newItems, toVisit{city, cost})
				newItems = append(newItems, after...)
				q.items = newItems
		}
	}
}
func (q *sortedQueue) takeHead() *toVisit {
	if len(q.items) == 0 {
		return nil
	}

	head := q.items[0]
	q.items = q.items[1:]

	return &head
}
func (q *sortedQueue) length() int {
	return len(q.items)
}
