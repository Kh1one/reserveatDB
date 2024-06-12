package entities

type OngoingMenu struct {
	ReservedOrderId int
	ReservationId   int
	MenuId          int
	MenuName        string
	Amount          int
	OrderPrice      int
}
