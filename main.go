package main

import (
	"crud-dasar/app"
	"net/http"
)

func main() {

	db := app.NewDB()

	server := http.Server{
		Addr: "localhost:3000",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}