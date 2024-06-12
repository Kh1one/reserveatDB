package employeecontroller

import (
	"errors"
	"log"
	"net/http"
	"test1/config"
	"test1/entities/employee"
	employeemodel "test1/models"

	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	PhoneNum string `validate:"required"`
	Password string `validate:"required"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("get")

		temp, err := template.ParseFiles("views/employee/login.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, nil)
		}
	}

	if r.Method == "POST" {
		r.ParseForm()

		//log.Println("post")
		userInput := &UserInput{
			PhoneNum: r.Form.Get("phoneNum"),
			Password: r.Form.Get("password"),
		}
		var message error

		var data employee.Employee = employeemodel.GetEmployeeLoginData(userInput.PhoneNum)

		if data.EmployeePhonenum != userInput.PhoneNum { //data isn't found
			log.Println("data not found")
			message = errors.New("incorrect phone number or password")

			tempData := map[string]interface{}{
				"error": message,
			}

			temp, err := template.ParseFiles("views/employee/login.html")
			if err != nil {
				panic(err.Error())
			} else {
				temp.Execute(w, tempData)
			}
		} else { //data is found
			log.Println("found")

			errPassword := bcrypt.CompareHashAndPassword([]byte(data.Pass), []byte(userInput.Password))

			if errPassword != nil {
				message = errors.New("incorrect phone number or password")
			}

			if message != nil {
				tempData := map[string]interface{}{
					"error": message,
				}

				temp, err := template.ParseFiles("views/employee/login.html")
				if err != nil {
					panic(err.Error())
				} else {
					temp.Execute(w, tempData)
				}
			} else {
				//set session

				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["phoneNum"] = data.EmployeePhonenum
				session.Values["name"] = data.EmployeeName
				session.Values["userID"] = data.EmployeeId
				session.Values["position"] = data.Position

				session.Save(r, w)

				if session.Values["position"] == "waiter" {
					//http.Redirect(w, r, "/adminhome", http.StatusSeeOther)

				} else if session.Values["position"] == "cashier" {
					http.Redirect(w, r, "/walkinorder", http.StatusSeeOther)

				} else if session.Values["position"] == "admin" {
					http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
				}

				//http.Redirect(w, r, "/employeehome", http.StatusSeeOther)
			}
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	userData := map[string]any{
		"name":     session.Values["name"],
		"position": session.Values["position"],
	}

	if session.Values["position"] == "waiter" {
		temp, _ := template.ParseFiles("views/waiter/waiterhome.html")
		temp.Execute(w, userData)

	} else if session.Values["position"] == "cashier" {
		temp, _ := template.ParseFiles("views/cashier/cashierhome.html")
		temp.Execute(w, userData)

	} else if session.Values["position"] == "admin" {
		http.Redirect(w, r, "/adminhome", http.StatusSeeOther)

	}

}
