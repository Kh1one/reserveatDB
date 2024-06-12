package entities

type WalkinTransactions struct {
	WalkinTransactionsId int
	WalkinId             int
	CashierId            int
	WaiterId             int
	TotalPrice           int
	PaymentMethod        string
}
