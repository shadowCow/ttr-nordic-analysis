package main

import (
	"fmt"
	"sort"

	"github.com/shadowcow/ttrn_analysis/analysis"
)

func main() {
	tickets := analysis.Tickets()

	counts := countTicketsByCity(tickets)

	sorted := sortByValueDescending(counts)
	for _, s := range sorted {
		fmt.Printf("%s: %v\n", s.Key, s.Value)
	}

	fmt.Println()
	fmt.Println()

	counts2 := countTicketsValueByCity(tickets)
	sorted2 := sortByValuesDescending(counts2)
	for _, s := range sorted2 {
		fmt.Printf("%s: %v, %v\n", s.Key, s.Value.Count, s.Value.TotalValue)
	}

	fmt.Println()

	routesByCity := analysis.RoutesByCity()
	for k, v := range routesByCity {
		fmt.Printf("%s: %v\n\n", k, v)
	}

	graph := analysis.NewGraph(analysis.Routes())
	shortestPaths := analysis.NewShortestPaths(graph)
	cityOccurencesOnTicketShortestPaths := map[analysis.City]int{}
	for _, c := range analysis.ListCities() {
		count := countCityOccurrenceOnTicketShortestPaths(c, tickets, &shortestPaths)
		cityOccurencesOnTicketShortestPaths[c] = count
	}
	sortedOccurrences := sortByValueDescending(cityOccurencesOnTicketShortestPaths)
	fmt.Println("==============================================")
	fmt.Println("City Occurrence Count on Ticket Shortest Paths")
	fmt.Println("==============================================")
	for _, s := range sortedOccurrences {
		fmt.Printf("%s: %v\n", s.Key, s.Value)
	}

	routeOccurrencesOnTicketShortestPaths := map[analysis.Route]int{}
	for _, r := range analysis.Routes() {
		count := countRouteOccurrenceOnTicketShortestPaths(r, tickets, &shortestPaths)
		routeOccurrencesOnTicketShortestPaths[r] = count
	}
	sortedRouteOccurrences := sortRoutesByValueDescending(routeOccurrencesOnTicketShortestPaths)
	fmt.Println("==============================================")
	fmt.Println("Route Occurrence Count on Ticket Shortest Paths")
	fmt.Println("==============================================")
	for _, s := range sortedRouteOccurrences {
		fmt.Printf("%v: %v\n", s.Key, s.Value)
	}
}

func countTicketsByCity(tickets []analysis.Ticket) map[analysis.City]int {
	counts := map[analysis.City]int{}

	for _, ticket := range tickets {
		fromCount, ok := counts[ticket.From]
		if ok {
			counts[ticket.From] = fromCount + 1
		} else {
			counts[ticket.From] = 1
		}

		toCount, ok := counts[ticket.To]
		if ok {
			counts[ticket.To] = toCount + 1
		} else {
			counts[ticket.To] = 1
		}
	}

	return counts
}

// Convert the map to a slice of key-value pairs (struct)
type keyValue struct {
	Key   analysis.City
	Value int
}

func sortByValueDescending(m map[analysis.City]int) []keyValue {
	// Convert the map to a slice of key-value pairs (struct)
    var keyValueSlice []keyValue

    for key, value := range m {
        keyValueSlice = append(keyValueSlice, keyValue{Key: key, Value: value})
    }

    // Define a custom sorting function for descending order
    sort.Slice(keyValueSlice, func(i, j int) bool {
        return keyValueSlice[i].Value > keyValueSlice[j].Value
    })

	return keyValueSlice
}

func sortRoutesByValueDescending(m map[analysis.Route]int) []routeKeyValue {
	// Convert the map to a slice of key-value pairs (struct)
    var keyValueSlice []routeKeyValue

    for key, value := range m {
        keyValueSlice = append(keyValueSlice, routeKeyValue{Key: key, Value: value})
    }

    // Define a custom sorting function for descending order
    sort.Slice(keyValueSlice, func(i, j int) bool {
        return keyValueSlice[i].Value > keyValueSlice[j].Value
    })

	return keyValueSlice
}

type routeKeyValue struct {
	Key analysis.Route
	Value int
}


type TicketCountValue struct {
	Count int
	TotalValue int
}

func countTicketsValueByCity(tickets []analysis.Ticket) map[analysis.City]TicketCountValue {
	counts := map[analysis.City]TicketCountValue{}

	for _, ticket := range tickets {
		fromValues, ok := counts[ticket.From]
		if ok {
			counts[ticket.From] = TicketCountValue{
				Count: fromValues.Count + 1,
				TotalValue: fromValues.TotalValue + ticket.Value,
			}
		} else {
			counts[ticket.From] = TicketCountValue{
				Count: 1,
				TotalValue: ticket.Value,
			}
		}

		toValues, ok := counts[ticket.To]
		if ok {
			counts[ticket.To] = TicketCountValue{
				Count: toValues.Count + 1,
				TotalValue: toValues.TotalValue + ticket.Value,
			}
		} else {
			counts[ticket.To] = TicketCountValue{
				Count: 1,
				TotalValue: ticket.Value,
			}
		}
	}

	return counts
}

// Convert the map to a slice of key-value pairs (struct)
type keyValues struct {
	Key   analysis.City
	Value TicketCountValue
}

func sortByValuesDescending(m map[analysis.City]TicketCountValue) []keyValues {
	// Convert the map to a slice of key-value pairs (struct)
    var keyValueSlice []keyValues

    for key, value := range m {
        keyValueSlice = append(keyValueSlice, keyValues{Key: key, Value: value})
    }

    // Define a custom sorting function for descending order
    sort.Slice(keyValueSlice, func(i, j int) bool {
        return keyValueSlice[i].Value.TotalValue > keyValueSlice[j].Value.TotalValue
    })

	return keyValueSlice
}

func countCityOccurrenceOnTicketShortestPaths(
	city analysis.City,
	tickets []analysis.Ticket,
	shortestPaths *analysis.ShortestPaths,
) int {
	count := 0
	for _, ticket := range tickets {
		_, path := shortestPaths.ShortestPath(ticket.From, ticket.To)
		fmt.Printf("ticket: %v\npath: %v\n\n", ticket, path)
		for _, pathCity := range path {
			if pathCity == city {
				count += 1
			}
		}
	}

	return count
}

func countRouteOccurrenceOnTicketShortestPaths(
	route analysis.Route,
	tickets []analysis.Ticket,
	shortestPaths *analysis.ShortestPaths,
) int {
	count := 0
	for _, ticket := range tickets {
		_, path := shortestPaths.ShortestPath(ticket.From, ticket.To)
		// fmt.Printf("ticket: %v\npath: %v\n\n", ticket, path)
		if pathHasRoute(path, route) {
			count += 1
		}
	}

	return count
}

func pathHasRoute(path []analysis.City, route analysis.Route) bool {
	for i, c := range path {
		if i < len(path) - 1 {
			from := c
			to := path[i + 1]

			if from == route.From && to == route.To || from == route.To && to == route.From {
				return true
			}
		}
	}

	return false
}
