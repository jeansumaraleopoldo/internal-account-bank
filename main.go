package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/jeansumaraleopoldo/internal-account-bank/controllers"
)

const (
	dbUrl  = "DB_URL"
	dbName = "DB_NAME"
)

func init() {
}

func main() {
	r := controller.Route()

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
