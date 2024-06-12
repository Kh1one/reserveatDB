package models

import (
	"log"
	"test1/config"
	"test1/entities"
)

func GetMenuDetails(menuId int) entities.Menu {
	row := config.DB.QueryRow("SELECT * FROM menu WHERE Menu_ID = ?", menuId)

	var data entities.Menu
	err := row.Scan(&data.MenuId, &data.MenuName, &data.Price, &data.MenuCategory)
	log.Println(data.MenuId, data.MenuName, data.Price, data.MenuCategory)

	if err != nil {
		return data
	} else {
		return data
	}
}

func SeatAvailableCheck(seatId int, date string, time string) int {
	row := config.DB.QueryRow("SELECT Reservation_ID FROM reservation WHERE Table_ID = ? AND Reservation_Date = ? AND Reservation_Time = ?", seatId, date, time)
	log.Println(seatId)
	log.Println(time)
	log.Println(date)

	var data int
	_ = row.Scan(&data)
	log.Println(data)

	return data
}

// buat masukkin info reservasi ke tabel reservation
func InsertReservation(tempReservation entities.Reservation) int {
	_, err := config.DB.Exec("INSERT INTO reservation (Customer_ID, Table_ID, Reservation_Date, Reservation_Time, Notes, Reservation_Status) VALUE (?, ?, ?, ?, ?, ?)", tempReservation.CustomerId, tempReservation.TableId, tempReservation.ReservationDate, tempReservation.ReservationTime, tempReservation.Notes, tempReservation.ReservationStatus)

	if err != nil {
		panic(err.Error())
	} else {
		result := config.DB.QueryRow("SELECT MAX(Reservation_ID) from reservation")

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

// buat masukkin menu yg diorder onlen
func InsertPreorder(item entities.ReservedOrder) {
	_, err := config.DB.Exec("INSERT INTO reservedorders (Reservation_ID, Menu_ID, Amount, Order_Price) VALUE (?, ?, ?, ?)", item.ReservationId, item.MenuId, item.Amount, item.OrderPrice)

	if err != nil {
		panic(err.Error())
	}
}

func InsertTransactions(id int, total int) {
	_, err := config.DB.Exec("INSERT INTO reservedtransactions (Reservation_ID, Total_Price, Payment_Method) VALUE (?, ?, ?)", id, total, "Qris")

	if err != nil {
		panic(err.Error())
	}
}

func GetOccupants(tableId int) int {
	row := config.DB.QueryRow("SELECT Capacity FROM seat WHERE Table_ID = ?", tableId)

	var data int
	_ = row.Scan(&data)
	log.Println(data)

	return data
}
