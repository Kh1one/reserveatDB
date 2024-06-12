package authcontroller

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"test1/config"
	"test1/entities"
	"test1/entities/employee"
	employeemodel "test1/models"
	usermodel "test1/models"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	PhoneNum string `validate:"required"`
	Password string `validate:"required"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("get")

		temp, err := template.ParseFiles("views/auth/login.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, nil)
		}
	}

	if r.Method == "POST" {
		r.ParseForm()

		log.Println("post")
		userInput := &UserInput{
			PhoneNum: r.Form.Get("phoneNum"),
			Password: r.Form.Get("password"),
		}
		var message error

		if userInput.PhoneNum != "" && userInput.Password != "" {
			var data entities.Customer = usermodel.GetLoginData(userInput.PhoneNum)

			if data.PhoneNum != userInput.PhoneNum { //data isn't found
				log.Println("data not found")

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
				} else if data.EmployeePhonenum == userInput.PhoneNum { //data is found
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
							http.Redirect(w, r, "/viewreservation", http.StatusSeeOther)
						} else if session.Values["position"] == "cashier" {
							http.Redirect(w, r, "/walkinorder", http.StatusSeeOther)

						} else if session.Values["position"] == "admin" {
							http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
						}

						//http.Redirect(w, r, "/employeehome", http.StatusSeeOther)
					}
				} else {
					message = errors.New("incorrect phone number or password")

					tempData := map[string]interface{}{
						"error": message,
					}

					temp, err := template.ParseFiles("views/auth/login.html")
					if err != nil {
						panic(err.Error())
					} else {
						temp.Execute(w, tempData)
					}
				}

			} else if data.PhoneNum == userInput.PhoneNum { //data is found
				log.Println("found")

				errPassword := bcrypt.CompareHashAndPassword([]byte(data.Pass), []byte(userInput.Password))

				if errPassword != nil {
					message = errors.New("incorrect phone number or password")
				}

				if message != nil {
					tempData := map[string]interface{}{
						"error": message,
					}

					temp, err := template.ParseFiles("views/auth/login.html")
					if err != nil {
						panic(err.Error())
					} else {
						temp.Execute(w, tempData)
					}
				} else {
					//a customer, not employee
					//set session

					session, _ := config.Store.Get(r, config.SESSION_ID)

					session.Values["loggedIn"] = true
					session.Values["phoneNum"] = data.PhoneNum
					session.Values["name"] = data.CustomerName
					session.Values["userID"] = data.CustomerId
					session.Values["position"] = "customer"

					session.Save(r, w)

					http.Redirect(w, r, "/home", http.StatusSeeOther)

				}
			}
		}

	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/auth/register.html")
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

		if user.Pass != cPassword || cPassword == "" {
			errorMessages["cpassword"] = "Password not confirmed"

		}

		if usermodel.CheckAvailable(user.PhoneNum) != 0 {

			//data is already in the database
			errorMessages["unavailable"] = "Phone number already has an account"
		}

		if len(errorMessages) > 0 {

			data := map[string]interface{}{
				"error":            1,
				"cpassError":       errorMessages["cpassword"],
				"userName":         user.CustomerName,
				"userPhonenum":     user.PhoneNum,
				"userPass":         user.Pass,
				"userCpass":        cPassword,
				"unavailableError": errorMessages["unavailable"],
			}

			temp, _ := template.ParseFiles("views/auth/register.html")
			temp.Execute(w, data)
		} else {

			// hashPassword
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
			user.Pass = string(hashPassword)

			// insert ke database
			usermodel.InsertUserData(user)

			http.Redirect(w, r, "/", http.StatusSeeOther)

		}

	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	//delete session
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		userData := map[string]any{
			"name":     session.Values["name"],
			"phoneNum": session.Values["phoneNum"],
		}

		temp, err := template.ParseFiles("views/home/profile.html")
		if err != nil {
			panic(err.Error())
		} else {
			temp.Execute(w, userData)
		}
	}

}
