// main.go
package main

import (

	"log"
	"net/http"
	"otp-auth-system/handlers"
	"otp-auth-system/utils"
    "github.com/joho/godotenv"
    

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }


    db, err := utils.InitDB() 
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    router := mux.NewRouter()
    handlers.InitializeRoutes(router, db)

    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
