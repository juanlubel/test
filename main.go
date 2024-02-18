package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"salu2/src/orm"
	"salu2/src/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	err = orm.Connect("mongodb://root:example@localhost:27017")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getenv("URL_BASE"))
	serverConfig := server.NewConfig("8080")
	server.Start(serverConfig)
}
