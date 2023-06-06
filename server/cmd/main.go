package main

import (
	"fmt"
	"os"
	"server/db"
	"server/internal/ws"
	"server/internal/user"
	"server/router"
	"server/utils"

	"github.com/joho/godotenv"
)

func main() {
	utils.ClearScreen()

	err := godotenv.Load()

	if err != nil {
		utils.Logger(4, "main.go -> godotenv.Load() -> Gagal membaca file .env")
	}

	dbug := os.Getenv("DEBUG_MODE")
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	fmt.Printf("\n > Server running on: %s", utils.Serv.Sprintf("http://%s:%s", host, port))
	fmt.Printf("\n\n > Route API:")
	fmt.Printf("\n   - /signup")
	fmt.Printf("\n   - /login")
	fmt.Printf("\n   - /logout")
	fmt.Printf("\n   - /ws/create_room")
	fmt.Printf("\n   - /ws/join_room/{:roomId}")
	fmt.Printf("\n   - /ws/get_rooms")
	fmt.Printf("\n   - /ws/get_clients/{:roomId}")
	fmt.Println("\n\n -------------------- [ LOG History ] --------------------\n")

	if dbug == "1" {
		utils.Logger(3, "main.go -> godotenv.Load()")
	}

	dbConn, err := db.ConnectDB()

	if err != nil {
		utils.Logger(4, err.Error())
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	address := fmt.Sprintf("%s:%s", host, port)

	router.InitRouter(userHandler, wsHandler)
	router.Start(address)
}
