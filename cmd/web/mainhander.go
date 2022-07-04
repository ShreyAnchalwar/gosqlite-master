package main

import (
	"encoding/json"
	"gosqlite/cmd/web/models"
	"gosqlite/internal/entity"
	"net/http"
	"time"
)

func (app *App) AddData(w http.ResponseWriter, r *http.Request) {

	var request models.AddData
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := entity.User{Key: request.Key, Value: request.Value, Timestamp: time.Now()}

	result := app.UserModel.Db.Create(&user) // pass pointer of data to create
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	jData, err := json.Marshal(&user)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

// For Get route

func (app *App) GetData(w http.ResponseWriter, r *http.Request) {

	var users []entity.User

	//destination struct is passed in
	app.UserModel.Db.Find(&users)

	jData, err := json.Marshal(&users)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}
