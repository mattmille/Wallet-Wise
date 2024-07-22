package main

import (
	"database/sql"
	"log"
	"wallet-wise/routes"

	"github.com/volatiletech/sqlboiler/boil"
	_ "modernc.org/sqlite"
)

func main() {

	// Open db
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		log.Fatal("Error: Cannot open sqlite3 database: ", err)
	}
	defer db.Close()

	boil.SetDB(db)

	//ctx := context.Background()
	//_, err = models.Expense().All(ctx, db)
	//if err != nil {
	//	log.Fatal("Error: Cannot open table categories: ", err)
	//}

	routes.Run()
}
