package homecontroller

import (
	"html/template"
	"log"
	"net/http"
	"test1/config"
	"test1/entities"
	homemodel "test1/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if session.Values["position"] == "waiter" {
		http.Redirect(w, r, "/viewreservation", http.StatusSeeOther)
	} else if session.Values["position"] == "cashier" {
		http.Redirect(w, r, "/walkinorder", http.StatusSeeOther)
	} else if session.Values["position"] == "admin" {
		http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
	} else {
		temp, _ := template.ParseFiles("views/home/index.html")

		temp.Execute(w, nil)
	}

}

func Menu(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		log.Println("post")

		session, _ := config.Store.Get(r, config.SESSION_ID)
		if len(session.Values) == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			r.ParseForm()

			menuId := r.FormValue("menuId")
			log.Println(menuId)

			http.Redirect(w, r, "/reservation", http.StatusSeeOther)
		}
	}

	temp, _ := template.ParseFiles("views/home/menu.html")
	temp.Execute(w, nil)
}

func Ongoing(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		userId := session.Values["userID"].(int)
		ongoing := homemodel.GetOngoing(userId)

		data := map[string]any{
			"data": ongoing,
		}

		temp, err := template.ParseFiles("views/home/ongoing.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, data)
		}
	}

}

func History(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "/history", http.StatusSeeOther)

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		history := homemodel.GetHistory(session.Values["userID"].(int))
		log.Println(history)

		var cancelledHistory []entities.History
		var OngoingHistory []entities.History
		var completedHistory []entities.History

		for _, data := range history {
			if data.ReservationStatus == "Cancelled" {
				cancelledHistory = append(cancelledHistory, data)
			} else if data.ReservationStatus == "Ongoing" {
				OngoingHistory = append(OngoingHistory, data)
			} else if data.ReservationStatus == "Completed" {
				completedHistory = append(completedHistory, data)
			}
		}

		data := map[string]any{
			"dataCompleted": completedHistory,
			"dataOngoing":   OngoingHistory,
			"dataCancelled": cancelledHistory,
		}

		temp, err := template.ParseFiles("views/home/history.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, data)
		}
	}
}
