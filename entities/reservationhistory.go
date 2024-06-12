package entities

type ReservationHistory struct {
	ReservedTransactionsId int
	ReservationId          int
	CashierId              int
	WaiterId               int
	TotalPrice             int
	PaymentMethod          string
	TableId                int
	ReservationDate        string
	ReservationTime        string
	Notes                  string
}
