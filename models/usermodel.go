package models

import (
	"log"
	"test1/config"
	"test1/entities"
)

func GetLoginData(phoneNum string) entities.Customer {
	row := config.DB.QueryRow("SELECT * FROM customer WHERE PhoneNum  = ?", phoneNum)

	var data entities.Customer
	row.Scan(&data.CustomerId, &data.CustomerName, &data.PhoneNum, &data.Pass)
	log.Println(data.CustomerId, data.CustomerName, data.PhoneNum, data.Pass)

	return data
}

func InsertUserData(user entities.Customer) {
	_, err := config.DB.Exec("insert into customer (Customer_Name, PhoneNum, Pass) values(?,?,?)", user.CustomerName, user.PhoneNum, user.Pass)

	if err != nil {
		panic(err.Error())
	}
}

func CheckAvailable(phoneNum string) int {
	row := config.DB.QueryRow("SELECT Customer_ID FROM customer WHERE PhoneNum  = ?", phoneNum)
	log.Println("available?")

	var data int
	row.Scan(&data)

	return data
}
