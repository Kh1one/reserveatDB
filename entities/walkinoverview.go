package entities

type WalkinOverview struct {
	WalkinId        int
	TransactionTime string
	TransactionDate string
	CashierId       int
	WaiterId        int
	TotalPrice      int
	PaymentMethod   string
}
