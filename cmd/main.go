package main

import (
	"flag"
	"github.com/StCredZero/rsrvtrst_thing/db"
	"github.com/StCredZero/rsrvtrst_thing/fib"
	"log"
)

type FibApp struct {
	dbconnect string
	db *db.DB
	fib *fib.Fibber
}

func (app *FibApp) initApp() {

	// cmd args
	initFlag := flag.Bool("init", false, "init db schema")
	dbconnect := flag.String("dbconnect", "postgres://postgres@localhost/fib_db?sslmode=disable", "specify dbconnection string")

	flag.Parse()

	if *initFlag {

	}

	app.dbconnect = *dbconnect
	app.db = new(db.DB)
	app.db.InitDatabase(*dbconnect)

	app.fib = fib.NewFibber()
}

func main() {

	app := new(FibApp)
	app.initApp()
	initEvApp(&evApp)
	go evApp.server.serverLoop()
	initWebserver(&evApp)
}