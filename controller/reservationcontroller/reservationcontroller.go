package reservationcontroller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test1/config"
	"test1/entities"
	reservationmodel "test1/models"
	"time"
)

type cartItem struct {
	MenuId      int
	MenuName    string
	SinglePrice int
	TotalPrice  int
	Amount      int
	Counter     int
}

var Counter int = 0

var cartItems = []cartItem{}

var cartTotal int

var totalItems int

var tempReservation entities.Reservation

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if session.Values["name"] == nil {
		log.Println("sdgv")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if r.Method == "GET" && session.Values["name"] != nil {
		log.Println("info get")

		temp, err := template.ParseFiles("views/reservation/info.html")
		if err != nil {
			panic(err.Error())
		} else {
			if tempReservation.ReservationDate != "" {

				data := map[string]any{
					"data": tempReservation,
				}
				temp.Execute(w, data)

			} else {
				temp.Execute(w, nil)

			}
		}
	}

	if r.Method == "POST" && session.Values["name"] != nil {
		log.Println("info post")
		r.ParseForm()

		date := r.Form.Get("date")
		tempReservation.ReservationDate = date
		log.Println(tempReservation.ReservationDate)

		time := r.Form.Get("time")
		tempReservation.ReservationTime = time
		log.Println(tempReservation.ReservationTime)

		notes := r.Form.Get("notes")
		tempReservation.Notes = notes
		log.Println(tempReservation.Notes)

		http.Redirect(w, r, "/reservation/seat", http.StatusSeeOther)

	}
}

func Seat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("seat get")

		temp, _ := template.ParseFiles("views/reservation/seat.html")

		temp.Execute(w, nil)

	}

	if r.Method == "POST" {
		log.Println("seat post")
		r.ParseForm()

		seat := r.Form.Get("seatid")
		seatId, err := strconv.Atoi(seat)
		log.Println(seatId)

		action := r.Form.Get("action")
		log.Println(action)

		if err != nil {
			panic(err.Error())
		} else {

			if reservationmodel.SeatAvailableCheck(seatId, tempReservation.ReservationDate, tempReservation.ReservationTime) == 0 {
				tempReservation.TableId = seatId
				log.Println(tempReservation.TableId)

				if action == "preorder" {
					http.Redirect(w, r, "/reservation/seat/menu", http.StatusSeeOther)

				} else if action == "skip" {
					http.Redirect(w, r, "/confirm ", http.StatusSeeOther)

				}

			} else {
				log.Println("someone's there")
				http.Redirect(w, r, "/reservation/seat", http.StatusSeeOther)
			}

		}
	}
}

func Menucns(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("views/reservation/menu.html")

	temp.Execute(w, nil)

}

func Menu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//show in cart

		cartTotal = 0
		totalItems = 0

		for i := range cartItems {
			cartTotal += cartItems[i].TotalPrice
			totalItems++

		}

		data := map[string]any{
			"totalItems": totalItems,
			"data":       cartItems,
			"cartTotal":  cartTotal,
		}

		//show in cart end
		log.Println("menu get")

		temp, err := template.ParseFiles("views/reservation/menu.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, data)
		}
	}

	if r.Method == "POST" {
		log.Println("menu post")
		r.ParseForm()

		action := r.Form.Get("action")

		if action == "decrease" || action == "increase" {
			var editMenuIndex string = r.FormValue("menuindex")
			var action string = r.FormValue("action")

			log.Println("POST count " + editMenuIndex)
			log.Println("POST tipe " + action)

			index, err := strconv.Atoi(editMenuIndex)
			if err != nil {
				panic(err.Error())
			} else {
				index--
				if action == "increase" {
					log.Println("increase")
					cartItems[index].Amount += 1

					cartItems[index].TotalPrice = cartItems[index].Amount * cartItems[index].SinglePrice

				} else if action == "decrease" {
					log.Println("decrease")

					if cartItems[index].Amount <= 1 {
						for i := range cartItems {
							if i >= index {
								cartItems[i].Counter--
							}
						}
						cartItems = append(cartItems[:index], cartItems[index+1:]...)
						Counter--
					} else {
						cartItems[index].Amount -= 1
						cartItems[index].TotalPrice = cartItems[index].Amount * cartItems[index].SinglePrice
					}
				}
			}

		} else {

			menu := r.Form.Get("menuId")
			log.Println("menu: ", menu)

			var amountNum = 1

			menuId, err := strconv.Atoi(menu)

			log.Println("id menu:", menuId)

			var flag int
			var index int

			for _, data := range cartItems {
				if menuId == data.MenuId {
					flag = 1
					index = data.Counter
				}
			}

			if flag == 1 {
				index--
				log.Println("increase")
				cartItems[index].Amount += 1

				cartItems[index].TotalPrice = cartItems[index].Amount * cartItems[index].SinglePrice
			} else {

				Counter++

				if err != nil {
					panic(err.Error())
				} else {

					var items entities.Menu = reservationmodel.GetMenuDetails(menuId)

					log.Println(items)

					cartItems = append(cartItems, cartItem{items.MenuId, items.MenuName, items.Price, amountNum * items.Price, amountNum, Counter})

				}
			}
		}
		//show in cart

		cartTotal = 0
		totalItems = 0

		for i := range cartItems {
			cartTotal += cartItems[i].TotalPrice
			totalItems++

		}

		data := map[string]any{
			"totalItems": totalItems,
			"data":       cartItems,
			"cartTotal":  cartTotal,
		}

		//show in cart end

		temp, _ := template.ParseFiles("views/reservation/menu.html")
		if cartTotal == 0 {
			temp.Execute(w, nil)
		} else {
			temp.Execute(w, data)
		}
	}

}

func Confirm(w http.ResponseWriter, r *http.Request) {
	var preOrder string

	if cartTotal == 0 {
		preOrder = "No"
	} else {
		preOrder = "Yes"
	}

	if r.Method == "GET" {

		log.Println(tempReservation)

		log.Println("confirm")
		session, _ := config.Store.Get(r, config.SESSION_ID)

		log.Println(preOrder)

		occupants := reservationmodel.GetOccupants(tempReservation.TableId)

		data := map[string]any{
			"cartData":  cartItems,
			"date":      tempReservation.ReservationDate,
			"time":      tempReservation.ReservationTime,
			"people":    occupants,
			"table":     tempReservation.TableId,
			"cartTotal": cartTotal,
			"name":      session.Values["name"].(string),
			"phoneNum":  session.Values["phoneNum"].(string),
			"preorder":  preOrder,
		}

		tempReservation.CustomerId = session.Values["userID"].(int)
		tempReservation.ReservationStatus = "Upcoming"

		if preOrder == "No" {
			temp, err := template.ParseFiles("views/reservation/confirm_no_preorder.html")
			if err != nil {
				panic(err.Error())
			} else {
				temp.Execute(w, data)
				var _ = reservationmodel.InsertReservation(tempReservation)
			}
		} else {
			temp, err := template.ParseFiles("views/reservation/confirm_preorder.html")
			if err != nil {
				panic(err.Error())
			} else {
				temp.Execute(w, data)
			}
		}
	}

}

func Payment(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("views/reservation/payment.html")
	data := map[string]any{
		"totalPayment": cartTotal,
	}
	temp.Execute(w, data)

}

func Complete(w http.ResponseWriter, r *http.Request) {

	var reservationID = reservationmodel.InsertReservation(tempReservation)

	var item entities.ReservedOrder

	//putting data into database
	for i := range cartItems {
		item.MenuId = cartItems[i].MenuId
		item.Amount = cartItems[i].Amount
		item.OrderPrice = cartItems[i].TotalPrice
		// item.OrderStatus = "Confirmed"
		item.ReservationId = reservationID

		reservationmodel.InsertPreorder(item)
	}
	reservationmodel.InsertTransactions(item.ReservationId, cartTotal)

	//resetting data
	Counter = 0
	cartTotal = 0
	cartItems = nil

	var DateTime = time.Now()
	str := strings.Split(DateTime.String(), " ")
	transactionDate := str[0]
	log.Println(transactionDate) //this is date

	str2 := strings.Split(str[1], ".")
	transactionTime := str2[0]
	log.Println(transactionTime) //this is time

	session, _ := config.Store.Get(r, config.SESSION_ID)

	data := map[string]any{
		"name": session.Values["name"],
		"time": transactionTime,
		"date": transactionDate,
	}

	temp, err := template.ParseFiles("views/reservation/complete.html")
	if err != nil {
		panic(err.Error())
	} else {
		temp.Execute(w, data)
	}
}
