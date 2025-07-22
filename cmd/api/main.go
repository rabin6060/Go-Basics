package main

import (
	"log"
	"github.com/rabin6060/Go-Basics/internal/db"
	"github.com/rabin6060/Go-Basics/internal/env"
	"github.com/rabin6060/Go-Basics/internal/store"

)

func main(){
	cfg := config{
			addr: env.GetString("Addr","=:8001"),
			db:dbConfig{
				addr: env.GetString("addr","postgres://postgres:postgres@localhost/social?sslmode=disable"),
				maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS",30),
				maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS",30),
				maxIdleTime: env.GetString("DB_MAX_IDLE_TIME","15m"),
			},
	}
	 
	db,err:=db.New(cfg.db.addr,cfg.db.maxOpenConns,cfg.db.maxIdleConns,cfg.db.maxIdleTime)
	if err!=nil{
		log.Panic(err)
	}
	defer db.Close()
	log.Println("database connected")
	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

	
}
