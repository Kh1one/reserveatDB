package models

import (
	"log"
	"test1/config"
	"test1/entities"
	"test1/entities/employee"
)

func GetEmployeeLoginData(phoneNum string) employee.Employee {
	admin := config.DB.QueryRow("SELECT * FROM admin WHERE PhoneNum  = ?", phoneNum)
	cashier := config.DB.QueryRow("SELECT * FROM cashier WHERE PhoneNum  = ?", phoneNum)
	waiter := config.DB.QueryRow("SELECT * FROM waiter WHERE PhoneNum  = ?", phoneNum)

	var data employee.Employee

	admin.Scan(&data.EmployeeId, &data.EmployeeName, &data.EmployeePhonenum, &data.Pass)
	log.Println("admin", admin)
	if data.EmployeeName == "" {
		//not admin
		cashier.Scan(&data.EmployeeId, &data.EmployeeName, &data.EmployeePhonenum, &data.Pass)
		log.Println("cashier", cashier)

		if data.EmployeeName == "" {
			//not cashier
			waiter.Scan(&data.EmployeeId, &data.EmployeeName, &data.EmployeePhonenum, &data.Pass)
			log.Println("waiter", waiter)

			if data.EmployeeName == "" {
				//not an employee
				data.Position = "none"
				return data
			}

			data.Position = "waiter"
			return data
		}

		data.Position = "cashier"
		return data
	}

	data.Position = "admin"
	return data
}

func CheckAvailableEmployee(phoneNum string) int {
	rowCashier := config.DB.QueryRow("SELECT Cashier_ID FROM cashier WHERE PhoneNum  = ?", phoneNum)
	rowWaiter := config.DB.QueryRow("SELECT Waiter_ID FROM waiter WHERE PhoneNum  = ?", phoneNum)
	rowAdmin := config.DB.QueryRow("SELECT Admin_ID FROM admin WHERE PhoneNum  = ?", phoneNum)

	var dataCashier int
	var dataWaiter int
	var dataAdmin int

	rowCashier.Scan(&dataCashier)
	rowWaiter.Scan(&dataWaiter)
	rowAdmin.Scan(&dataAdmin)

	if dataCashier != 0 || dataWaiter != 0 || dataAdmin != 0 {
		return 1
	} else {
		return 0
	}
}

func InsertEmployeeData(user entities.Customer, position string) {

	if position == "Waiter" {
		_, err := config.DB.Exec("insert into waiter (Waiter_Name, PhoneNum, Pass) values(?,?,?)", user.CustomerName, user.PhoneNum, user.Pass)

		if err != nil {
			panic(err.Error())
		}
	} else if position == "Cashier" {
		_, err := config.DB.Exec("insert into cashier (Cashier_Name, PhoneNum, Pass) values(?,?,?)", user.CustomerName, user.PhoneNum, user.Pass)

		if err != nil {
			panic(err.Error())
		}
	} else if position == "Admin" {
		_, err := config.DB.Exec("insert into admin (Admin_Name, PhoneNum, Pass) values(?,?,?)", user.CustomerName, user.PhoneNum, user.Pass)

		if err != nil {
			panic(err.Error())
		}
	}
}

func InsertWalkin(seat int, transactionDate string, transactionTime string) int {
	_, err := config.DB.Exec("INSERT INTO walkin (Table_ID, Transaction_Date, Transaction_Time) VALUE (?, ?, ?)", seat, transactionDate, transactionTime)

	if err != nil {
		panic(err.Error())
	} else {
		result := config.DB.QueryRow("SELECT MAX(WalkIn_ID) from walkin")

		var id int
		err := result.Scan(&id)
		log.Println(id)
		if err != nil {
			panic(err.Error())
		} else {
			return id
		}
	}
}

func InsertOrders(item entities.WalkinOrder) {
	_, err := config.DB.Exec("INSERT INTO walkinorders (WalkIn_ID, Menu_ID, Order_Price, Amount) VALUE (?, ?, ?, ?)", item.WalkinId, item.MenuId, item.OrderPrice, item.Amount)

	if err != nil {
		panic(err.Error())
	}
}

func InsertWalkinTransaction(walkinId int, cashier int, waiter string, totalPrice int, paymentMethod string) {
	_, err := config.DB.Exec("INSERT INTO walkintransactions (WalkIn_ID, Cashier_ID, Waiter_ID, Total_Price, Payment_Method) VALUE (?, ?, ?, ?, ?)", walkinId, cashier, waiter, totalPrice, paymentMethod)

	if err != nil {
		panic(err.Error())
	}
}

func GetReservation() []entities.ReservationInfo {
	data, err := config.DB.Query("SELECT reservation.Reservation_ID, customer.Customer_Name, reservation.Reservation_Date, reservation.Reservation_Time, reservation.Table_ID, reservation.Notes, reservation.Reservation_Status FROM reservation JOIN customer ON reservation.Customer_ID = customer.Customer_ID")

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.ReservationInfo

		for data.Next() {
			var temp entities.ReservationInfo
			err := data.Scan(&temp.ReservationId, &temp.CustomerName, &temp.Date, &temp.Time, &temp.Seat, &temp.Notes, &temp.Status)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func ChangeReservationStatus(newStatus string, id int) {
	_, err := config.DB.Exec("UPDATE reservation SET Reservation_Status = ? WHERE Reservation_ID = ?", newStatus, id)

	if err != nil {
		panic(err.Error())
	}
}

func GetReservedOrder(id int) []entities.OngoingMenu {
	data, err := config.DB.Query("SELECT reservedorders.ReservedOrder_ID, reservedorders.Reservation_ID, reservedorders.Menu_ID, menu.Menu_Name , reservedorders.Amount, reservedorders.Order_Price FROM reservedorders JOIN menu ON reservedorders.Menu_ID = menu.Menu_ID WHERE reservedorders.Reservation_ID = ?", id)

	if err != nil {
		panic(err.Error())
	} else {

		defer data.Close()

		var allData []entities.OngoingMenu

		for data.Next() {
			var temp entities.OngoingMenu
			err := data.Scan(&temp.ReservedOrderId, &temp.ReservationId, &temp.MenuId, &temp.MenuName, &temp.Amount, &temp.OrderPrice)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}
