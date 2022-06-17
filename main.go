package main

import (
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/server"
)

func main() {
	database.StartDB() // starta o banco de dados
	s := server.NewServer() // starta o server

	s.Run() 
}
