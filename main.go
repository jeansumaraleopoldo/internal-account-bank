package main

import (
	"fmt"
	"log"
	"net/http"

	controller "internal-account-bank/controllers"
)

func main() {
	r := controller.Route()

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
