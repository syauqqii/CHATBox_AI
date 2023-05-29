package main

import(
	"log"
	// "github.com/ADEVASATRIA/CHATBox_AI/server/db"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
func main() {
	_, err := DB.NewDatabase()
	if err != nil {
		log.Fatal("could not initialize database connection: %s", err)
	}

}
