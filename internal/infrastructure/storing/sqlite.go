package storing

import (
	"TodoApp_basic/routes/logger"
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

func NewSqlite(dns string, logger logger.Logger) (db *sql.DB, err error) {

	//dns = "./TodoApp.db" //filepath.Join("./", "test","db")

	// Create a new database with the desired schema
	db, err = sql.Open("sqlite", dns)
	if err != nil {
		logger.WithError(err)
		log.Fatal(err)
	}
	//defer db.Close()
	if err = createDatabase(db, logger); err != nil {
		logger.Error("Error creating data storing")
	}

	fmt.Println("Database created successfully!")
	logger.Info("Database created successfully!")
	return db, err
}

// createDatabase creates the necessary tables and schema for the database.
func createDatabase(db *sql.DB, logger logger.Logger) error {
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
		logger.WithError(err)
		return err
	}

	// Add more tables or schema changes as needed.

	return nil
}
