package entity

// Route value object
type Route struct {
	AirlineCode string // AirlineCode is the airline two digit code
	Origin      string // Origin airport IATA 3 code
	Destination string // Destination airport IATA 3 code
}
