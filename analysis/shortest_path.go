package analysis

import (
	"math"
	"slices"
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

// store the costs and paths for all shortest paths
// from any city to any other city
type ShortestPaths struct {
	costs map[City]map[City]int
	paths map[City]map[City][]City
}

func NewShortestPaths(g *Graph) ShortestPaths {
	s := ShortestPaths{
		costs: make(map[City]map[City]int),
		paths: make(map[City]map[City][]City),
	}
	for from := range g.Edges {
		shortestPathCosts, shortestPaths := dijkstra(g, from)
		
		s.costs[from] = shortestPathCosts
		s.paths[from] = prependFromCity(shortestPaths, from)

		// fmt.Printf("from: %v\nshortestPaths %v\n\n", from, s.paths[from])
	}

	return s
}

func prependFromCity(paths map[City][]City, from City) map[City][]City {
	withPrepends := make(map[City][]City, len(paths))
	for k, v := range paths {
		withCity := make([]City, len(v) + 1)
		withCity[0] = from
		for i, c := range v {
			withCity[i+1] = c
		}

		withPrepends[k] = withCity
	}

	return withPrepends
}

// returns the cost of the shortest path
// and the path through all intermediate cities.
func (s *ShortestPaths) ShortestPath(from, to City) (int, []City) {
	return s.costs[from][to], s.paths[from][to]
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
				subpaths[neighbor] = append([]City{}, subpaths[current.city]...)
				subpaths[neighbor] = append(subpaths[neighbor], neighbor)

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
	if slices.ContainsFunc(q.items, func (e toVisit) bool {
		return e.city == city
	}) {
		return
	}

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
