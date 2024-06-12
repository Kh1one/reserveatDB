package entities

type History struct {
	ReservationId     int
	TotalPrice        int
	PaymentMethod     string
	ReservationDate   string
	ReservationTime   string
	TableId           int
	Capacity          int
	ReservationStatus string
}
