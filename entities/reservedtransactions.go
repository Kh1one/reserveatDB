package entities

type ReservedTransactions struct {
	ReservedTransactionsId int
	ReservationId          int
	CashierId              int
	WaiterId               int
	TotalPrice             int
	PaymentMethod          string
}
