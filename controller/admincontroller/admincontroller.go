package admincontroller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test1/entities"
	adminmodel "test1/models"
	employeemodel "test1/models"

	"time"

	"golang.org/x/crypto/bcrypt"
)

var walkinOverviewInRange []entities.WalkinOverview
var reservedOverviewInRange []entities.ReservedOverview
var reservedMenuSalesInRange []entities.MenuSales
var walkinMenuSalesInRange []entities.MenuSales
var startDate string
var endDate string
var menuId int

func Home(w http.ResponseWriter, r *http.Request) {
	var currentDate string

	var DateTime = time.Now()
	str := strings.Split(DateTime.String(), " ")
	currentDate = str[0]
	log.Println("home", currentDate) //this is date

	walkinOrderOverview := adminmodel.GetTodaysWalkinTransactions(currentDate)
	reservedOrderOverview := adminmodel.GetTodaysReservedTransactions(currentDate)

	temp, err := template.ParseFiles("views/admin/adminhome.html")
	if err != nil {
		panic(err.Error())
	} else {
		if walkinOrderOverview != nil {
			log.Println("data")

			data := map[string]any{
				"walkinOverview":   walkinOrderOverview,
				"reservedOverview": reservedOrderOverview,
			}
			temp.Execute(w, data)

		} else {
			log.Println("nil")
			temp.Execute(w, nil)

		}
	}
}

func ViewReservedOrder(w http.ResponseWriter, r *http.Request) {

}

func ViewWalkinOrder(w http.ResponseWriter, r *http.Request) {

}

func ViewOverviewInRange(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/admin/overviewinrange.html")
		if err != nil {
			panic(err.Error())
		} else {
			if walkinOverviewInRange != nil {
				log.Println("data")

				data := map[string]any{

					"startDate":        startDate,
					"endDate":          endDate,
					"walkinOverview":   walkinOverviewInRange,
					"reservedOverview": reservedOverviewInRange,
				}
				temp.Execute(w, data)

			} else {
				log.Println("nil")
				temp.Execute(w, nil)

			}
		}
	}

	if r.Method == "POST" {
		r.ParseForm()

		startDate = r.FormValue("start")
		endDate = r.FormValue("end")

		walkinOverviewInRange = adminmodel.GetWalkinSalesInRange(startDate, endDate)
		reservedOverviewInRange = adminmodel.GetReservedSalesInRange(startDate, endDate)

		http.Redirect(w, r, "/overviewinrange", http.StatusSeeOther)

	}
}

func ViewMenuSales(w http.ResponseWriter, r *http.Request) {
	log.Println("menu sales")
	if r.Method == "GET" {
		log.Println("menu sales get")

		temp, err := template.ParseFiles("views/admin/menusales.html")
		if err != nil {
			panic(err.Error())
		} else {
			if reservedMenuSalesInRange != nil || walkinMenuSalesInRange != nil {

				data := map[string]any{
					"menuId":        menuId,
					"startDate":     startDate,
					"endDate":       endDate,
					"walkinSales":   reservedMenuSalesInRange,
					"reservedSales": walkinMenuSalesInRange,
				}
				log.Println("sales", data)

				temp.Execute(w, data)

			} else {
				log.Println("sales nil")
				temp.Execute(w, nil)

			}
		}
	}

	if r.Method == "POST" {
		r.ParseForm()

		startDate = r.FormValue("start")
		endDate = r.FormValue("end")
		menuIdString := r.FormValue("menuId")
		menuId, _ = strconv.Atoi(menuIdString)

		reservedMenuSalesInRange = adminmodel.GetReservedMenuSalesInRange(startDate, endDate, menuId)
		walkinMenuSalesInRange = adminmodel.GetWalkinMenuSalesInRange(startDate, endDate, menuId)

		http.Redirect(w, r, "/menusales", http.StatusSeeOther)

	}
}

func RegisterEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/admin/registeremployee.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, nil)
		}

	} else if r.Method == "POST" {
		// mengambil inputan form
		r.ParseForm()

		user := entities.Customer{
			CustomerName: r.Form.Get("name"),
			PhoneNum:     r.Form.Get("phonenum"),
			Pass:         r.Form.Get("pass"),
		}

		errorMessages := make(map[string]interface{})

		var cPassword = r.Form.Get("cpass")
		var position = r.Form.Get("position")

		if user.CustomerName == "" {
			errorMessages["name"] = "Name field cannot be empty"
			log.Println(errorMessages["name"])
		}
		if position == "" || (position != "Waiter" && position != "Cashier" && position != "Admin") {
			errorMessages["position"] = "Please fill position field correctly"
			log.Println(errorMessages["position"])
		}
		if user.PhoneNum == "" {
			errorMessages["phonenum"] = "Phone number field cannot be empty"
		}
		if user.Pass == "" {
			errorMessages["password"] = "Password field cannot be empty"
		}
		if user.Pass != cPassword || cPassword == "" {
			errorMessages["cpassword"] = "Password not confirmed"
		}

		if employeemodel.CheckAvailableEmployee(user.PhoneNum) != 0 {
			log.Println("whops")
			//data is already in the database
			errorMessages["unavailable"] = "Phone number already has an account"
		}

		if len(errorMessages) > 0 {

			data := map[string]interface{}{
				"error":            1,
				"nameError":        errorMessages["name"],
				"phonenumError":    errorMessages["phonenum"],
				"passError":        errorMessages["password"],
				"cpassError":       errorMessages["cpassword"],
				"positionError":    errorMessages["position"],
				"position":         position,
				"userName":         user.CustomerName,
				"userPhonenum":     user.PhoneNum,
				"userPass":         user.Pass,
				"userCpass":        cPassword,
				"unavailableError": errorMessages["unavailable"],
			}

			temp, _ := template.ParseFiles("views/admin/registeremployee.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
			user.Pass = string(hashPassword)

			// insert ke database
			employeemodel.InsertEmployeeData(user, position)
			data := map[string]interface{}{
				"message": "Registration success!",
			}
			temp, _ := template.ParseFiles("views/admin/registeremployee.html")
			temp.Execute(w, data)

		}

	}

}
