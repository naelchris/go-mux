package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //imported pq for postgresql
	"log"
)

type App struct{
	Router *mux.Router
	DB *sql.DB
}


func (a *App) Initialize(user, password, dbname string){
	//initialize method will take in the details required to connect to the database.
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string){
	//run method will start the application
}