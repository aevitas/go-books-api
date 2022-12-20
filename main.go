package main

import "aevitas.dev/go-books/api"

func main() {
	srv := &api.Server{}

	srv.InitDb("host=localhost user=kiruna password=lol dbname=books")
	srv.InitGin()

	srv.RegisterRoutes()

	srv.Start(":8050")
}
