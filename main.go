package main

import(
	"log"
	// "github.com/ADEVASATRIA/CHATBox_AI/server/DB"
	"chatbox_ai/server/db"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

}
