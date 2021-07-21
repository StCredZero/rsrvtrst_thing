package main

import (
	"flag"
	"github.com/StCredZero/rsrvtrst_thing/db"
	"github.com/StCredZero/rsrvtrst_thing/fib"
	"github.com/StCredZero/rsrvtrst_thing/webserver"
)

type FibApp struct {
	dbconnect string
	db *db.DB
	fib *fib.Fibber
	srv *webserver.Server
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
	app.fib.Initialize()

	app.srv = new(webserver.Server)
	app.srv.Fibber = app.fib
}

func main() {
	app := new(FibApp)
	app.initApp()
	app.srv.Start()
}