package database

import (
	"os"

	"github.com/go-bongo/bongo"
)

func connect() *bongo.Connection {
	var config *bongo.Config

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		config = &bongo.Config{
			ConnectionString: dbURL,
		}
	} else {
		config = &bongo.Config{
			ConnectionString: "localhost:27017/testing",
		}
	}

	db, err := bongo.Connect(config)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

var db = connect()
