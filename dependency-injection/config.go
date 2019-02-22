package main

import (
	"database/sql"
)

//Enabled tells us if our application should return real data.
//DatabasePath tells us where our database lives (we’re using sqlite).
//Port tells us the port on which we’ll be running our server.
type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}

//open our database connection. It relies on our Config and returns a *sql.DB
func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}
