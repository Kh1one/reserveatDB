package main

import (
	"log"
	"net/http"
	"test1/config"
	"test1/controller/admincontroller"
	"test1/controller/authcontroller"
	cashiercontroller "test1/controller/cashiercontroller"
	"test1/controller/employeecontroller"
	"test1/controller/homecontroller"
	"test1/controller/reservationcontroller"
	"test1/controller/waitercontroller"
)

func main() {
	config.ConnectDB()

	//homepage
	http.HandleFunc("/", homecontroller.Home)
	http.HandleFunc("/menu", homecontroller.Menu)

	//customer
	http.HandleFunc("/ongoing", homecontroller.Ongoing)
	http.HandleFunc("/history", homecontroller.History)

	//login customer
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/profile", authcontroller.Profile)

	//reservation
	http.HandleFunc("/reservation", reservationcontroller.Index)
	http.HandleFunc("/reservation/seat", reservationcontroller.Seat)
	http.HandleFunc("/reservation/seat/menu", reservationcontroller.Menu)
	http.HandleFunc("/confirm", reservationcontroller.Confirm)
	http.HandleFunc("/payment", reservationcontroller.Payment)
	http.HandleFunc("/complete", reservationcontroller.Complete)

	//login employee
	http.HandleFunc("/employeelogin", employeecontroller.Login)
	http.HandleFunc("/employeehome", employeecontroller.Home)

	//cashier
	http.HandleFunc("/walkinorder", cashiercontroller.Order)
	http.HandleFunc("/walkinorder/confirm", cashiercontroller.Confirm)

	//waiter
	http.HandleFunc("/viewreservation", waitercontroller.ViewReservation)
	http.HandleFunc("/viewreservedmenu", waitercontroller.ViewReservedMenu)

	//admin
	http.HandleFunc("/adminhome", admincontroller.Home)
	http.HandleFunc("/overviewinrange", admincontroller.ViewOverviewInRange)
	http.HandleFunc("/menusales", admincontroller.ViewMenuSales)
	http.HandleFunc("/registeremployee", admincontroller.RegisterEmployee)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	log.Println(("server running"))
	http.ListenAndServe(":5050", nil)
}
