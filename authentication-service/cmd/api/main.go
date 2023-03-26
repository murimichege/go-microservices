package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8081"
type Config struct{
	DB *sql.DB
	Models data.Models

}
func main(){
	log.Println(("Starting authentication service"))
	 // toDo Connect to database

	 // set up config

	 app:=Config{}
	 svr := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	 }
	 
	 err := svr.ListenAndServe()
	 if err != nil {
		log.Panic(err)
	 }

}