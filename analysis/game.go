package analysis

import "fmt"


type City string

const (
	Murmansk City = "Murmansk"
	Kirkenes City =	"Kirkenes"
	Honningsvag City =	"Honningsvag"
	Tromso City =	"Tromso"
	Narvik City =	"Narvik"
	Kiruna City =	"Kiruna"
	Boden City =	"Boden"
	Tornio City =	"Tornio"
	Rovaniemi City =	"Rovaniemi"
	Oulu City =	"Oulu"
	Kajaani City =	"Kajaani"
	Lieksa City =	"Lieksa"
	Kuopio City =	"Kuopio"
	Imatra City =	"Imatra"
	Helsinki City =	"Helsinki"
	Lahti City =	"Lahti"
	Turku City ="Turku"
	Tampere City =	"Tampere"
	Vaasa City = "Vaasa"
	Umea City = "Umea"
	Sundsvall City = "Sundsvall"
	Tallinn City ="Tallinn"
	Stockholm City ="Stockholm"
	Norrkoping City = "Norrkoping"
	Karlskrona City = "Karlskrona"
	Kobenhavn City = "Kobenhavn"
	Arhus City = "Arhus"
	Alborg City = "Alborg"
	Goteborg City = "Goteborg"
	Oslo City = "Oslo"
	Orebro City = "Orebro"
	Kristiansand City = "Kristiansand"
	Stavangar City = "Stavangar"
	Bergen City = "Bergen"
	Andalsnes City = "Andalsnes"
	Lillehammer City = "Lillehammer"
	Trondheim City = "Trondheim"
	Ostersund City = "Ostersund"
	Moirana City = "Moirana"
)

type Ticket struct {
	From City
	To City
	Value int
}

type Route struct {
	From City
	To City
	Length int
}

var routes []Route = []Route{
	{Oslo, Orebro, 2},
	{Oslo, Goteborg, 2},
	{Oslo, Alborg, 3},
	{Oslo, Kristiansand, 2},
	{Oslo, Bergen, 4},
	{Oslo, Lillehammer, 2},
	{Bergen, Stavangar, 2},
	{Bergen, Andalsnes, 5},
	{Stavangar, Kristiansand, 2},
	{Stavangar, Kristiansand, 3},
	{Kristiansand, Alborg, 2},
	{Alborg, Arhus, 1},
	{Alborg, Goteborg, 2},
	{Arhus, Kobenhavn, 1},
	{Kobenhavn, Goteborg, 2},
	{Kobenhavn, Karlskrona, 2},
	{Goteborg, Orebro, 2},
	{Goteborg, Norrkoping, 3},
	{Karlskrona, Norrkoping, 3},
	{Norrkoping, Stockholm, 1},
	{Norrkoping, Orebro, 2},
	{Orebro, Sundsvall, 4},
	{Orebro, Stockholm, 2},
	{Stockholm, Tallinn, 4},
	{Stockholm, Helsinki, 4},
	{Stockholm, Turku, 3},
	{Stockholm, Sundsvall, 4},
	{Helsinki, Tallinn, 2},
	{Helsinki, Turku, 1},
	{Helsinki, Imatra, 3},
	{Helsinki, Lahti, 1},
	{Helsinki, Tampere, 1},
	{Imatra, Lahti, 2},
	{Imatra, Kuopio, 2},
	{Turku, Tampere, 1},
	{Tampere, Lahti, 1},
	{Lahti, Kuopio, 3},
	{Kuopio, Lieksa, 1},
	{Kuopio, Kajaani, 2},
	{Kuopio, Vaasa, 4},
	{Kuopio, Oulu, 3},
	{Tampere, Vaasa, 2},
	{Vaasa, Sundsvall, 3},
	{Vaasa, Umea, 1},
	{Vaasa, Oulu, 3},
	{Kajaani, Lieksa, 1},
	{Kajaani, Oulu, 2},
	{Lieksa, Murmansk, 9},
	{Murmansk, Kirkenes, 3},
	{Kirkenes, Rovaniemi, 5},
	{Kirkenes, Honningsvag, 2},
	{Honningsvag, Tromso, 4},
	{Tromso, Narvik, 3},
	{Narvik, Kiruna, 1},
	{Narvik, Moirana, 4},
	{Kiruna, Boden, 3},
	{Boden, Tornio, 1},
	{Boden, Umea, 3},
	{Tornio, Rovaniemi, 1},
	{Tornio, Oulu, 1},
	{Rovaniemi, Oulu, 2},
	{Umea, Sundsvall, 3},
	{Sundsvall, Ostersund, 2},
	{Ostersund, Trondheim, 2},
	{Trondheim, Moirana, 5},
	{Trondheim, Moirana, 6},
	{Trondheim, Lillehammer, 3},
	{Trondheim, Andalsnes, 2},
	{Andalsnes, Lillehammer, 2},
}

var tickets []Ticket = []Ticket{
	{Helsinki, Kiruna, 10},
	{Helsinki, Kirkenes, 13},
	{Helsinki, Ostersund, 8},
	{Kobenhavn, Murmansk, 24},
	{Stockholm, Tromso, 17},
	{Oslo, Honningsvag, 21},
	{Arhus, Lillehammer, 6},
	{Alborg, Norrkoping, 5},
	{Goteborg, Andalsnes, 6},
	{Tromso, Vaasa, 11},
	{Stavangar, Karlskrona, 8},
	{Tampere, Tallinn, 3},
	{Alborg, Umea, 11},
	{Narvik, Murmansk, 12},
	{Norrkoping, Boden, 11},
	{Sundsvall, Lahti, 6},
	{Orebro, Kuopio, 10},
	{Narvik, Tallinn, 13},
	{Tampere, Boden, 6},
	{Helsinki, Lieksa, 5},
	{Stavangar, Rovaniemi, 18},
	{Oslo, Stockholm, 4},
	{Oslo, Vaasa, 9},
	{Oslo, Stavangar, 4},
	{Oslo, Helsinki, 8},
	{Oslo, Moirana, 10}, 
	{Oslo, Kobenhavn, 4},
	{Stockholm, Kajaani, 10},
	{Stockholm, Umea, 7},
	{Stockholm, Bergen, 8},
	{Stockholm, Imatra, 7},
	{Stockholm, Kobenhavn, 6},
	{Kobenhavn, Oulu, 14},
	{Bergen, Kobenhavn, 8},
	{Bergen, Trondheim, 7},
	{Helsinki, Bergen, 12},
	{Helsinki, Kobenhavn, 10},
	{Kobenhavn, Narvik, 18},
	{Bergen, Tornio, 17},
	{Bergen, Narvik, 16},
	{Goteborg, Oulu, 12},
	{Turku, Trondheim, 10},
	{Goteborg, Turku, 7},
	{Tornio, Imatra, 6},
	{Tampere, Kristiansand, 10},
	{Kristiansand, Moirana, 12},
}

func Routes() []Route {
	return routes
}

func Tickets() []Ticket {
	return tickets
}

func RoutesByCity() map[City][]Route {
	routesByCity := map[City][]Route{}

	for _, route := range routes {
		if _, exists := routesByCity[route.From]; !exists {
			routesByCity[route.From] = []Route{route}
		} else {
			routesByCity[route.From] = append(routesByCity[route.From], route)
		}

		if _, exists := routesByCity[route.To]; !exists {
			routesByCity[route.To] = []Route{route}
		} else {
			routesByCity[route.To] = append(routesByCity[route.To], route)
		}
	}

	return routesByCity
}

func PointsForRouteLength(length int) int {
	switch length {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 7
	case 5:
		return 10
	case 6:
		return 15
	case 9:
		return 27
	default:
		panic(fmt.Sprintf("illegal route length: %v", length))
	}
}