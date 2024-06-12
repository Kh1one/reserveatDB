package entities

type ReservedOrder struct {
	ReservedOrderId int
	ReservationId   int
	MenuId          int
	Amount          int
	OrderPrice      int
}
