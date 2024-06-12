package models

import (
	"test1/config"
	"test1/entities"
)

func GetTodaysWalkinTransactions(currentDate string) []entities.WalkinOverview {
	data, err := config.DB.Query("SELECT walkin.Walkin_ID, walkintransactions.Waiter_ID, walkintransactions.Cashier_ID, walkin.Transaction_Time, walkintransactions.Total_Price, walkintransactions.Payment_Method FROM walkin JOIN walkintransactions ON walkin.walkin_ID = walkintransactions.walkin_ID WHERE walkin.Transaction_Date = ?", currentDate)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.WalkinOverview

		for data.Next() {
			var temp entities.WalkinOverview
			err := data.Scan(&temp.WalkinId, &temp.WaiterId, &temp.CashierId, &temp.TransactionTime, &temp.TotalPrice, &temp.PaymentMethod)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func GetTodaysReservedTransactions(currentDate string) []entities.ReservedOverview {
	data, err := config.DB.Query("SELECT reservation.Reservation_ID, reservation.Reservation_Time, reservedtransactions.Total_Price, reservedtransactions.Payment_Method FROM reservation JOIN reservedtransactions ON reservation.Reservation_ID = reservedtransactions.Reservation_ID WHERE reservation.Reservation_Date = ?", currentDate)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.ReservedOverview

		for data.Next() {
			var temp entities.ReservedOverview
			err := data.Scan(&temp.ReservedId, &temp.TransactionTime, &temp.TotalPrice, &temp.PaymentMethod)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func GetReservedSalesInRange(startDate string, endDate string) []entities.ReservedOverview {
	data, err := config.DB.Query("SELECT reservation.Reservation_ID, reservation.Reservation_Date, reservation.Reservation_Time, reservedtransactions.Total_Price, reservedtransactions.Payment_Method FROM reservation JOIN reservedtransactions ON reservation.Reservation_ID = reservedtransactions.Reservation_ID WHERE reservation.Reservation_Date >= ? AND reservation.Reservation_Date <= ?", startDate, endDate)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.ReservedOverview

		for data.Next() {
			var temp entities.ReservedOverview
			err := data.Scan(&temp.ReservedId, &temp.TransactionDate, &temp.TransactionTime, &temp.TotalPrice, &temp.PaymentMethod)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func GetWalkinSalesInRange(startDate string, endDate string) []entities.WalkinOverview {
	data, err := config.DB.Query("SELECT walkin.Walkin_ID, walkintransactions.Waiter_ID, walkintransactions.Cashier_ID, walkin.Transaction_Date, walkin.Transaction_Time, walkintransactions.Total_Price, walkintransactions.Payment_Method FROM walkin JOIN walkintransactions ON walkin.walkin_ID = walkintransactions.walkin_ID WHERE walkin.Transaction_Date >= ? AND walkin.Transaction_Date <= ?", startDate, endDate)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.WalkinOverview

		for data.Next() {
			var temp entities.WalkinOverview
			err := data.Scan(&temp.WalkinId, &temp.WaiterId, &temp.CashierId, &temp.TransactionDate, &temp.TransactionTime, &temp.TotalPrice, &temp.PaymentMethod)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func GetWalkinMenuSalesInRange(startDate string, endDate string, menuId int) []entities.MenuSales {
	data, err := config.DB.Query("SELECT walkin.WalkIn_ID, walkin.Transaction_Date, walkinorders.Amount, walkinorders.Order_Price FROM walkin JOIN walkinorders ON walkin.WalkIn_ID = walkinorders.WalkinOrders_ID WHERE walkin.Transaction_Date > ? AND walkin.Transaction_Date < ? AND walkinorders.Menu_ID = ?", startDate, endDate, menuId)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.MenuSales

		for data.Next() {
			var temp entities.MenuSales
			err := data.Scan(&temp.Id, &temp.TransactionDate, &temp.Amount, &temp.Price)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func GetReservedMenuSalesInRange(startDate string, endDate string, menuId int) []entities.MenuSales {
	data, err := config.DB.Query("SELECT reservation.Reservation_ID, reservation.Reservation_Date, reservedorders.Amount, reservedorders.Order_Price FROM reservation JOIN reservedorders ON reservation.Reservation_ID = reservedorders.Reservation_ID WHERE reservation.Reservation_Date > ? AND reservation.Reservation_Date < ? AND reservedorders.Menu_ID = ?", startDate, endDate, menuId)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.MenuSales

		for data.Next() {
			var temp entities.MenuSales
			err := data.Scan(&temp.Id, &temp.TransactionDate, &temp.Amount, &temp.Price)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}
