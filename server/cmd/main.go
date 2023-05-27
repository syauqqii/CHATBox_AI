package main

import(
	"log"
	"server/DB"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("could not initialize database connection: %s", err)
	}

}
