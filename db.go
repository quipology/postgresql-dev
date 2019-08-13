package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Actor is a reflection of the actor table within the dvdrental database
type Actor struct {
	ID         int
	FirstName  string
	LastName   string
	LastUpdate time.Time
}

const (
	host     = "localhost"
	port     = 5432
	dbname   = "dvdrental"
	user     = "postgres"
	password = "xxxxx"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	psqlInfo := fmt.Sprintf("sslmode=disable host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test database connectivity
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Connection to database successful!")
	}
	// Delete actor Bobby if exists
	_, err = db.Exec(`DELETE FROM actor WHERE first_name='Bobby' AND last_name='Williams';`)
	checkError(err)

	// ReAdd actor Bobby
	_, err = db.Exec(`INSERT INTO actor(actor_id, first_name, last_name, last_update) VALUES($1, $2, $3, $4);`, 39443, "Bobby", "Williams", time.Now())
	checkError(err)

	// Query all actors and look for record recently entered
	actors := []Actor{}

	rows, err := db.Query(`SELECT * FROM actor;`)
	for rows.Next() {
		var actor Actor
		err = rows.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.LastUpdate)
		if err != nil {
			fmt.Println("Something went wrong with that row:", err)
		}
		actors = append(actors, actor)
	}
	// If desired record found print the record to stdout
	for _, i := range actors {
		if i.FirstName == "Bobby" && i.LastName == "Williams" {
			fmt.Printf("Record for Bobby Williams found: %v", i)
		}
	}

}
