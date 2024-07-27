package main

import (
	"fmt"
	"log"
	"quote/data"
	"quote/servers"

	"github.com/joho/godotenv"
)

func main() {

	db, err := data.ConnectToDataBase()
	if err != nil {
		log.Fatal("DB connection is Fail!")
	}

	db.Run()

	envFile, _ := godotenv.Read(".env")
	listenAddr := envFile["SERVERADDRESS"]
	allowedIP := envFile["ALLOWEDIP"]

	server := servers.NewServer(listenAddr, db, allowedIP)
	fmt.Println("Server is running on port :", listenAddr)
	server.Start()
}
