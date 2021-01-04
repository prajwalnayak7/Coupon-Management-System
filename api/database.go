package api

import (
	"os" 
	"database/sql"
	"log"
	"time"
)

func ConnectToDB() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	conn, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:5000)/"+database)
	if err != nil {
		log.Println("Failed to connect to the Database")
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("Database", database, "connected successfully as user", user)
	}
	// Set the maximum number of concurrently open connections (in-use + idle)
	// to 5. Setting this to less than or equal to 0 will mean there is no 
	// maximum limit (which is also the default setting).
	conn.SetMaxOpenConns(5)
	// Set the maximum number of concurrently idle connections to 5. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(5*time.Minute)
	// defer conn.Close()
	DB = conn
}
