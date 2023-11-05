package main

import (
	"fmt"
	"sort"

	"github.com/shadowcow/ttrn_analysis/analysis"
)

func main() {
	// cities := []City{
	// 	"Murmansk",
	// 	"Kirkenes",
	// 	"Honningsvag",
	// 	"Tromso",
	// 	"Narvik",
	// 	"Kiruna",
	// 	"Boden",
	// 	"Tornio",
	// 	"Rovaniemi",
	// 	"Oulu",
	// 	"Kajaani",
	// 	"Lieksa",
	// 	"Kuopio",
	// 	"Imatra",
	// 	"Helsinki",
	// 	"Lahti",
	// 	"Turku",
	// 	"Tampere",
	// 	"Vaasa",
	// 	"Umea",
	// 	"Sundsvall",
	// 	"Tallinn",
	// 	"Stockholm",
	// 	"Norrkoping",
	// 	"Karlskrona",
	// 	"Kobenhavn",
	// 	"Arhus",
	// 	"Alborg",
	// 	"Goteborg",
	// 	"Oslo",
	// 	"Orebro",
	// 	"Kristiansand",
	// 	"Stavangar",
	// 	"Bergen",
	// 	"Andalsnes",
	// 	"Lillehammer",
	// 	"Trondheim",
	// 	"Ostersund",
	// 	"Moirana",
	// }

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