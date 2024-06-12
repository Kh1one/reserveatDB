package waitercontroller

import (
	"log"
	"net/http"
	"strconv"
	"test1/entities"
	employeemodel "test1/models"
	"text/template"
)

func ViewReservation(w http.ResponseWriter, r *http.Request) {

	resvInfo := employeemodel.GetReservation()

	var upcomingReservation = []entities.ReservationInfo{}
	var completedReservation = []entities.ReservationInfo{}
	var cancelledReservation = []entities.ReservationInfo{}
	var ongoingReservation = []entities.ReservationInfo{}

	for _, data := range resvInfo {
		if data.Status == "Upcoming" {
			upcomingReservation = append(upcomingReservation, data)
		} else if data.Status == "Completed" {
			completedReservation = append(completedReservation, data)
		} else if data.Status == "Cancelled" {
			cancelledReservation = append(cancelledReservation, data)
		} else if data.Status == "Ongoing" {
			ongoingReservation = append(ongoingReservation, data)
		}
	}

	if r.Method == "GET" {
		log.Println("view reservation get")

		temp, err := template.ParseFiles("views/waiter/viewreservation.html")
		if err != nil {
			panic(err.Error())
		} else {
			if resvInfo != nil {

				data := map[string]any{
					"upcomingData":  upcomingReservation,
					"completedData": completedReservation,
					"cancelledData": cancelledReservation,
					"ongoingData":   ongoingReservation,
				}
				temp.Execute(w, data)

			} else {
				temp.Execute(w, nil)

			}
		}
	}

	if r.Method == "POST" {
		log.Println("view reservation post")

		r.ParseForm()

		var action string = r.FormValue("action")
		var reservationId string = r.FormValue("reservationId")
		reservationIdNum, _ := strconv.Atoi(reservationId)

		if action == "complete" {
			employeemodel.ChangeReservationStatus("Completed", reservationIdNum)
		} else if action == "cancel" {
			employeemodel.ChangeReservationStatus("Cancelled", reservationIdNum)
		} else if action == "ongoing" {
			employeemodel.ChangeReservationStatus("Ongoing", reservationIdNum)
		}

		http.Redirect(w, r, "/viewreservation", http.StatusSeeOther)

	}
}

var ViewMenuReservationId int

func ViewReservedMenu(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var reservationId string = r.FormValue("reservationId")
	reservationIdNum, _ := strconv.Atoi(reservationId)

	if reservationId != "" {
		ViewMenuReservationId = reservationIdNum
	}
	reservedMenu := employeemodel.GetReservedOrder(ViewMenuReservationId)

	temp, err := template.ParseFiles("views/waiter/viewreservedmenu.html")
	if err != nil {
		panic(err.Error())
	} else {
		if reservedMenu != nil {
			log.Println("data")

			data := map[string]any{
				"reservationId":   reservedMenu[0].ReservationId,
				"confirmedOrders": reservedMenu,
			}
			temp.Execute(w, data)

		} else {
			log.Println("nil")
			temp.Execute(w, nil)

		}
	}

}
