package main

import (
	"log"
)

func main(){
	app := &application{
		config: config{
			addr: ":8001",
		},
	}

	log.Fatal(app.run())

	
}
