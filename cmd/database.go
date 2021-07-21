package main

import (
	"github.com/jmoiron/sqlx"
)

func initSchema(db *sqlx.DB) error {

	var schema = `
    	CREATE TABLE fibonaci (
            ordinal        bigint,
			value          bigint,
			CONSTRAINT fibonaci_pk PRIMARY KEY (ordinal)
        );
		CREATE UNIQUE INDEX ordinal_index ON fibonaci(ordinal);
		`
	_, err := db.Exec(schema)
	return err
}

type fibonaciEntry struct {
	Ordinal int64
	Value   int64
}

func createFibonaciEntry(ordinal, value uint64, db *sqlx.DB) error {
	f := &fibonaciEntry{
		Ordinal: int64(ordinal),
		Value:   int64(value),
	}
	_, err := db.NamedExec(
		`INSERT INTO fibonaci
		(ordinal, value)
		VALUES (:ordinal, :value)`,
		f,
	)
	return err
}

func updateFibonaciEntry(ordinal, value uint64, db *sqlx.DB) error {
	f := &fibonaciEntry{
		Ordinal: int64(ordinal),
		Value:   int64(value),
	}
	_, err := db.NamedExec(
		`UPDATE fibonaci SET
			ordinal=:ordinal,
			value=:value`,
		f,
	)
	return err
}