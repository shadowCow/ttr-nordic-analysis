package analysis


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

