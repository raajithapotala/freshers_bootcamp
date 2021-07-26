package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"freshers_bootcamp/day3/exercise2/Config"
	"freshers_bootcamp/day3/exercise2/Controllers"
	"freshers_bootcamp/day3/exercise2/Models"
	"freshers_bootcamp/day3/exercise2/Routes"

	"github.com/jinzhu/gorm"
)


func TestGet(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//Setting the router
	router := Routes.SetupRouter()
	router.GET("/user-api/user/", Controllers.GetUsers)

	//Get request
	request, _ := http.NewRequest("GET", "/user-api/user/", nil)

	//Recording the response
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `[{"id":1,"name":"Raajitha","LastName":"Potala","DOB":"22-03-2000","address":"Visakhapatnam","subject":"Communication","marks":78},{"id":2,"name":"Raajitha","LastName":"Potala","DOB":"22-03-2000","address":"Visakhapatnam","subject":"Electronics","marks":89},{"id":3,"name":"Raajitha","LastName":"Potala","DOB":"22-03-2000","address":"Visakhapatnam","subject":"Data Structures","marks":95},{"id":4,"name":"Shireen","LastName":"Meher","DOB":"29-06-2000","address":"Visakhapatnam","subject":"Data Structures","marks":95},{"id":5,"name":"Shireen","LastName":"Meher","DOB":"29-06-2000","address":"Visakhapatnam","subject":"Electronics","marks":85},{"id":6,"name":"Shireen","LastName":"Meher","DOB":"29-06-2000","address":"Visakhapatnam","subject":"Communication","marks":81},{"id":123,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":124,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":69},{"id":125,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":126,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":127,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":128,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":129,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":130,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":131,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":132,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":133,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":134,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":135,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85},{"id":136,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":137,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":138,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":139,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71},{"id":140,"name":"","LastName":"","DOB":"","address":"","subject":"","marks":71}]`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

func TestPost(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.User{})


	//Setting the router
	router := Routes.SetupRouter()
	router.POST("/user-api/user/",Controllers.CreateUser)

	//send request
	newStud := Models.User{
		Name: "Test",
		LastName: "Check",
		DOB: "xyz",
		Address: "location",
		Subject: "Maths",
		Marks: 85,
	}

	responseBody,_ := json.Marshal(newStud)
	req, _ := http.NewRequest("POST", "/user-api/user/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `{"id":144,"name":"Test","LastName":"Check","DOB":"xyz","address":"location","subject":"Maths","marks":85}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

func TestPut(t *testing.T) {
	//SQL database using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.User{})

	//setting up router
	router := Routes.SetupRouter()
	router.PUT("/user-api/user/1",Controllers.UpdateUser)

	//send request
	newStudent := Models.User{
		LastName: "P",
		Marks:     85,
	}

	responseBody,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("PUT", "/user-api/user/1/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	expectedOutput := `{"id":1,"name":"Raajitha","LastName":"Potala","DOB":"22-03-2000","address":"Visakhapatnam","subject":"Communication","marks":85}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}
}

func TestDelete(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.DELETE("/user-api/user/2/",Controllers.DeleteUser )

	//Get request
	req, _ := http.NewRequest("DELETE", "/user-api/user/21/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	//checking test case
	if response.Code != 307 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Code, 200)
	}
}