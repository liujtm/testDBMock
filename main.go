package main

import (
	"database/sql"
	"fmt"
	"log"
	"testDBMock/internal"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create users table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize service using Wire
	userService, err := internal.InitializeUserService(db)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new user
	err = userService.CreateUser("john_doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Get the user
	user, err := userService.GetUser(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User: ID=%d, Username=%s, Email=%s\n", user.ID, user.Username, user.Email)
}
