package models

import (
	"database/sql"
	"log"
	"test1/config"
	"test1/entities"
)

func GetOngoing(customerId int) []entities.History {
	rows, err := config.DB.Query("SELECT reservation.Reservation_ID, reservation.Reservation_Date, reservation.Reservation_Time, reservation.Table_ID, reservation.Reservation_Status, seat.Capacity FROM reservation JOIN seat ON reservation.Table_ID = seat.Table_ID WHERE Customer_ID = ? AND Reservation_Status = ?", customerId, "Upcoming")

	if err != nil {
		panic(err.Error())
	} else {

		defer rows.Close()

		var allData []entities.History

		for rows.Next() {
			var temp entities.History
			err := rows.Scan(&temp.ReservationId, &temp.ReservationDate, &temp.ReservationTime, &temp.TableId, &temp.ReservationStatus, &temp.Capacity)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}

func GetHistory(id int) []entities.History {
	rows, err := config.DB.Query("SELECT reservation.Reservation_ID, reservation.Reservation_Date, reservation.Reservation_Time, reservation.Table_ID, reservation.Reservation_Status, seat.Capacity FROM reservation JOIN seat ON reservation.Table_ID = seat.Table_ID WHERE Customer_ID = ? AND (reservation.Reservation_Status = ? OR reservation.Reservation_Status = ? OR reservation.Reservation_Status = ?);", id, "Completed", "Cancelled", "Ongoing")

	if err != nil {
		panic(err.Error())
	} else {

		defer rows.Close()
		var allData []entities.History

		var temp entities.History

		for rows.Next() {

			err := rows.Scan(&temp.ReservationId, &temp.ReservationDate, &temp.ReservationTime, &temp.TableId, &temp.ReservationStatus, &temp.Capacity)

			if err != nil {
				panic(err.Error())
			} else {
				log.Println(temp)

				row := config.DB.QueryRow("SELECT Total_Price, Payment_Method FROM reservedtransactions WHERE Reservation_ID = ?", temp.ReservationId)

				errr := row.Scan(&temp.TotalPrice, &temp.PaymentMethod)

				if errr != nil && errr != sql.ErrNoRows {
					temp.TotalPrice = 0
					temp.PaymentMethod = ""

				}
				allData = append(allData, temp)
				log.Println(temp)

			}
		}
		return allData
	}
}

func GetAllMenu() []entities.Menu {
	rows, err := config.DB.Query("SELECT * FROM menu")

	if err != nil {
		panic(err.Error())
	} else {

		defer rows.Close()

		var allData []entities.Menu

		for rows.Next() {
			var temp entities.Menu
			err := rows.Scan(&temp.MenuId, &temp.MenuName, &temp.Price, &temp.MenuCategory)

			if err != nil {
				panic(err.Error())
			} else {
				allData = append(allData, temp)
			}
		}
		return allData
	}
}
