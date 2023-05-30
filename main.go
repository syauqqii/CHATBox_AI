package main

import (
	"chatbox_ai/server/db"
	"chatbox_ai/server/user"
	"chatbox_ai/server/router"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)
	// userRep := user.NewUserRepository(dbConn.GetDB())
	// userSvc := user.NewService(userRep)
	// useHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")

}
