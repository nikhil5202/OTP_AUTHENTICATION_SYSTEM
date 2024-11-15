//handlers/routes.go
package handlers

import (
    "database/sql"
    "github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router, db *sql.DB) {
    router.HandleFunc("/register", RegisterUser(db)).Methods("POST")
    router.HandleFunc("/login", LoginUser(db)).Methods("POST")
    router.HandleFunc("/resend-otp", ResendOTP(db)).Methods("POST")
    router.HandleFunc("/user/details", GetUserDetails(db)).Methods("GET")
}
