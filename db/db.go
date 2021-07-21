package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	lib *sqlx.DB
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
	return err
}

type fibonaciEntry struct {
	Ordinal int64
	Value   int64
}

func (db *DB) CreateFibonaciEntry(ordinal, value uint64) error {
	f := &fibonaciEntry{
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

func (db *DB) UpdateFibonaciEntry(ordinal, value uint64) error {
	f := &fibonaciEntry{
		Ordinal: int64(ordinal),
		Value:   int64(value),
	}
	_, err := db.lib.NamedExec(
		`UPDATE fibonaci SET
			ordinal=:ordinal,
			value=:value`,
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