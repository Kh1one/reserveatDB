package entities

type Reservation struct {
	ReservationId     int
	CustomerId        int
	TableId           int
	ReservationDate   string
	ReservationTime   string
	ReservationStatus string
	Notes             string
}
