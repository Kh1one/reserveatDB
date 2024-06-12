package waitercontroller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test1/config"
	"test1/entities"
	employeemodel "test1/models"
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

var counter int
var tempOrder = []cartItem{}
var totalPrice int //as in cart's total price
var seatId int

func Order(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("walkin order get")

		temp, err := template.ParseFiles("views/cashier/walkin.html")
		if err != nil {
			panic(err.Error())
		} else {
			if tempOrder != nil {

				data := map[string]any{
					"data": tempOrder,
				}
				temp.Execute(w, data)

			} else {
				temp.Execute(w, nil)

			}
		}
	}

	if r.Method == "POST" {
		log.Println("walkin post")
		r.ParseForm()

		if r.Form.Get("action") == "addcart" {
			menu := r.Form.Get("menuid")
			menuId, _ := strconv.Atoi(menu)

			num := r.Form.Get("amount")
			amount, _ := strconv.Atoi(num)

			var flag int
			var i int

			for _, data := range tempOrder {
				if menuId == data.MenuId {
					flag = 1
					i = data.Counter
				}
			}

			if flag == 1 {
				i--
				tempOrder[i].Amount += amount
				tempOrder[i].TotalPrice = tempOrder[i].SinglePrice * tempOrder[i].Amount
			} else {
				counter++
				var tempMenu = reservationmodel.GetMenuDetails(menuId)
				tempOrder = append(tempOrder, cartItem{tempMenu.MenuId, tempMenu.MenuName, tempMenu.Price, tempMenu.Price * amount, amount, counter})
			}

		} else if r.Form.Get("action") == "increase" {
			var editMenuIndex string = r.FormValue("menuindex")
			index, _ := strconv.Atoi(editMenuIndex)
			index--

			log.Println("increase")
			tempOrder[index].Amount += 1

			tempOrder[index].TotalPrice = tempOrder[index].SinglePrice * tempOrder[index].Amount
		} else if r.Form.Get("action") == "decrease" {
			var editMenuIndex string = r.FormValue("menuindex")
			index, _ := strconv.Atoi(editMenuIndex)
			index--

			log.Println("decrease")

			if tempOrder[index].Amount <= 1 {

				for i := range tempOrder {
					if i >= index {
						tempOrder[i].Counter--
					}
				}

				tempOrder = append(tempOrder[:index], tempOrder[index+1:]...)
				counter--
			} else {

				tempOrder[index].Amount -= 1
				tempOrder[index].TotalPrice = tempOrder[index].SinglePrice * tempOrder[index].Amount
			}
		} else if r.Form.Get("action") == "confirm" {
			r.ParseForm()

			var seat = r.Form.Get("seat")
			temp, _ := strconv.Atoi(seat)
			seatId = temp
			http.Redirect(w, r, "/walkinorder/confirm", http.StatusSeeOther)

		}

		totalPrice = 0
		for i := range tempOrder {
			totalPrice += tempOrder[i].TotalPrice
		}

		data := map[string]any{
			"data":      tempOrder,
			"cartTotal": totalPrice,
		}

		temp, err := template.ParseFiles("views/cashier/walkin.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, data)
		}
	}
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	//payment methods
	var DateTime = time.Now()
	str := strings.Split(DateTime.String(), " ")
	transactionDate := str[0]
	log.Println(transactionDate) //this is date

	str2 := strings.Split(str[1], ".")
	transactionTime := str2[0]
	log.Println(transactionTime) //this is time

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if r.Method == "GET" {
		log.Println("walkin order get")

		temp, _ := template.ParseFiles("views/cashier/confirmwalkin.html")

		data := map[string]any{
			"cashier":   session.Values["userID"].(int),
			"date":      transactionDate,
			"time":      transactionTime,
			"table":     seatId,
			"cartData":  tempOrder,
			"cartTotal": totalPrice,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		r.ParseForm()

		var waiter string = r.FormValue("waitername")
		var paymentMethod string = r.FormValue("payment")

		var walkinId = employeemodel.InsertWalkin(seatId, transactionDate, transactionTime)

		var item entities.WalkinOrder

		for i := range tempOrder {
			item.WalkinId = walkinId
			item.MenuId = tempOrder[i].MenuId
			item.Amount = tempOrder[i].Amount
			item.OrderPrice = tempOrder[i].TotalPrice

			employeemodel.InsertOrders(item)
		}
		cashier := session.Values["userID"].(int)

		employeemodel.InsertWalkinTransaction(walkinId, cashier, waiter, totalPrice, paymentMethod)
		log.Println("date: " + transactionDate + "  time: " + transactionTime)

		temp, _ := template.ParseFiles("views/cashier/walkincomplete.html")

		data := map[string]any{
			"cashier":       session.Values["userID"].(int),
			"waiter":        waiter,
			"date":          transactionDate,
			"time":          transactionTime,
			"table":         seatId,
			"cartData":      tempOrder,
			"cartTotal":     totalPrice,
			"paymentMethod": paymentMethod,
		}

		counter = 0
		totalPrice = 0
		tempOrder = nil
		temp.Execute(w, data)
	}

}
