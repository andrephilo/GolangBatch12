package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sesi11/models"

	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request := models.LoginRequest{}
	response := models.LoginResponse{}
	reqbyte, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(reqbyte, &request); err != nil {
		fmt.Println("Request gagal")

		response.Code = "02"
		response.Message = "Request gagal"
	}
	fmt.Println(request.Username)
	fmt.Println(request.Password)
	fmt.Println(request.Age)
	fmt.Println(request.Gender)
	fmt.Println(request.Email)

	var regname = request.Username
	var regpassword = request.Password
	var regage = request.Age
	var reggender = request.Gender
	var regemail = request.Email

	db, err := sql.Open("mysql", "root:asdasdxx@/golang?charset=utf8")
	if err != nil {
		fmt.Println("Error connect DB =>", err)
		return
	}
	db.Ping()

	stmt, err := db.Prepare(`INSERT INTO test2 values(?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println("Error create stmt =>", err)
		return
	}

	if regname == "admin" && regpassword == "admin" {
		_, err := stmt.Exec(1, regname, regage, reggender, regemail, regpassword, "admin")
		if err != nil {
			fmt.Println("Error query execute =>", err)
			return
		}
		return
	} else {
		_, err := stmt.Exec(rand.Intn(10), regname, regage, reggender, regemail, regpassword, "user")
		if err != nil {
			fmt.Println("Error query execute =>", err)
			return
		}
	}

	if regname != "" && regpassword != "" && regage != 0 && reggender != "" && regemail != "" {
		response.Code = "00"
		response.Message = "Sukses"
		jsonRes, _ := json.Marshal(response)
		w.Write(jsonRes)
		return
	}

	response.Code = "01"
	response.Message = "Gagal"
	jsonRes, _ := json.Marshal(response)
	w.Write(jsonRes)
	return
}

func Result(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./views/result.gtpl")

	db, err := sql.Open("mysql", "root:asdasdxx@/golang?charset=utf8")
	if err != nil {
		fmt.Println("Error connect DB =>", err)
		return
	}
	db.Ping()

	rows, err := db.Query(`Select name, age, gender, email, password from test2`)
	if err != nil {
		fmt.Println("Error select data =>", err)
		return
	}

	type data struct {
		resname     string
		resage      int
		resgender   string
		resemail    string
		respassword string
	}

	for rows.Next() {
		var name string
		var age int
		var gender string
		var email string
		var password string
		rows.Scan(&name, &age, &gender, &email, &password)

		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(gender)
		fmt.Println(email)
		fmt.Println(password)

		view := data{}
		view.resname = name
		view.resage = age
		view.resgender = gender
		view.resemail = email
		view.respassword = password

		if err := t.Execute(w, view); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

}

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/login.gtpl")
	t.Execute(w, nil)
}
