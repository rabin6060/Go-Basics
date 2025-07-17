package main

import (
	"log"

	"github.com/rabin6060/Go-Basics/internal/env"
	"github.com/rabin6060/Go-Basics/internal/store"
)

func main(){

	store := store.NewStorage(nil) 
	app := &application{
		config: config{
			addr: env.GetString("Addr",":8001"),
		},
		store: store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))

	
}
