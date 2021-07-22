package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	lib *sqlx.DB
}

func (db *DB) DropTable() error {
	var sql = `
    	DROP TABLE IF EXISTS fibonaci;
		`
	_, err := db.lib.Exec(sql)
	return err
}

func (db *DB) InitSchema() error {
	var schema = `
    	CREATE TABLE fibonaci (
            ordinal        bigint,
			value          bigint,
			CONSTRAINT fibonaci_pk PRIMARY KEY (ordinal)
        );
		CREATE UNIQUE INDEX ordinal_index ON fibonaci(ordinal);
		`
	_, err := db.lib.Exec(schema)
	if err != nil {
		return err
	}
	err = db.CreateFibonaciEntry(0, 0)
	if err != nil {
		return err
	}
	err = db.CreateFibonaciEntry(1, 1)

	return err
}

type FibonaciEntry struct {
	Ordinal int64
	Value   int64
}

func (db *DB) CreateFibonaciEntry(ordinal, value uint64) error {
	f := &FibonaciEntry{
		Ordinal: int64(ordinal),
		Value:   int64(value),
	}
	_, err := db.lib.NamedExec(
		`INSERT INTO fibonaci
		(ordinal, value)
		VALUES (:ordinal, :value)`,
		f,
	)
	return err
}

func (db *DB) InitDatabase(dbconnect string) {
	var err error
	//LOG(dbconnect)
	db.lib, err = sqlx.Connect("postgres", dbconnect)
	if err != nil {
		panic("Opening DB: " + err.Error())
	}
}

func (db *DB) RetrieveData(fn func(ordinal, value uint64)) {
	rows, err := db.lib.Queryx("select ordinal, value from fibonaci")
	defer rows.Close()
	if err != nil {
		log.Fatal("RetrieveData: ", err)
	}
	var ordinal, value int64
	for rows.Next() {
		err = rows.Scan(&ordinal, &value)
		if err != nil {
			log.Fatal("RetrieveData-loop: ", err)
		}
		fn(uint64(ordinal), uint64(value))
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("RetrieveData-end: ", err)
	}
}
