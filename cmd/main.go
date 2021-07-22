package main

import (
	"flag"
	"github.com/StCredZero/rsrvtrst_thing/db"
	"github.com/StCredZero/rsrvtrst_thing/fib"
	"github.com/StCredZero/rsrvtrst_thing/webserver"
	"log"
	"os"
	"time"
)

type FibApp struct {
	dbconnect string
	db        *db.DB
	fib       *fib.Fibber
	srv       *webserver.Server
}

func (app *FibApp) initApp() {

	// cmd args
	initFlag := flag.Bool("init", false, "init db schema")
	timeFlag := flag.Int64("time", 0, "execute a time test")
	dbconnect := flag.String("dbconnect", "postgres://fib_app@localhost/fib_db?sslmode=disable&password=secret123", "specify dbconnection string")

	flag.Parse()

	app.dbconnect = *dbconnect
	app.db = new(db.DB)
	app.db.InitDatabase(*dbconnect)

	// this activates the functionality for the init flag, then quits
	if *initFlag {
		app.db.DropTable()
		app.db.InitSchema()
		log.Printf("DB Schema Initialized")
		os.Exit(0)
	}

	app.fib = fib.NewFibber(app.db)
	app.fib.RetrieveData()

	if *timeFlag > 0 {
		start := time.Now()
		log.Printf("Starting timing session for N=%d", *timeFlag)
		app.fib.SyncExtendToN(uint64(*timeFlag))
		elapsed := time.Since(start)
		log.Printf("SyncExtendToN took %s", elapsed)
		os.Exit(0)
	}

	app.srv = new(webserver.Server)
	app.srv.Fibber = app.fib
}

func main() {
	app := new(FibApp)
	app.initApp()
	app.srv.Start()
}
