package tests

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func TestCleanUp(dns string, db *sql.DB) {

	err := db.Close()
	if err != nil {
		log.Print("error closing db")
	}

	err = os.Remove(dns)
	if err != nil {
		log.Print("error removing directory")
	}
}

func TestInit() (dsn string, db *sql.DB, err error) {

	dsn = "./test.db" //filepath.Join("./", "test","db")

	// Check if the database file exists
	if _, err := os.Stat(dsn); err == nil {
		// Database file exists, so delete it
		if err = os.Remove(dsn); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Deleted existing database")
	}

	// Create a new database with the desired schema
	db, err = sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	if err = createDatabase(db); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database created successfully!")
	return dsn, db, err
}

// createDatabase creates the necessary tables and schema for the database.
func createDatabase(db *sql.DB) error {
	// Create your table(s) here. This is just an example.
	_, err := db.Exec(`
		CREATE TABLE tasks (
			id CHARACTER ,
			createdAt INTEGER,
			updatedAt INTEGER,
			item VARCHAR(50),
			done INTEGER,
			doneAt INTEGER
		)
	`)
	if err != nil {
		return err
	}

	// Add more tables or schema changes as needed.

	return nil
}
