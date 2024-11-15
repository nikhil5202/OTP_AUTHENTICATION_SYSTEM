// services/user.go
package services

import (
	"database/sql"
	"log"
	"otp-auth-system/models"
)

func CreateUser(db *sql.DB, user models.User) error {
    log.Printf(user.MobileNumber,user.Name)
    _, err := db.Exec("INSERT INTO users (mobile_number, name) VALUES (?, ?)", user.MobileNumber, user.Name)
    return err
}

func GetUserDetails(db *sql.DB, userID int) (*models.User, error) {
    var user models.User
    query := "SELECT id, mobile_number, name FROM users WHERE id = ?"
    err := db.QueryRow(query, userID).Scan(&user.ID, &user.MobileNumber, &user.Name)
    return &user, err
}

func GetUserByMobile(db *sql.DB, mobile string) (*models.User, error) {
    var user models.User
    query := "SELECT id, mobile_number, name FROM users WHERE mobile_number = ?"
    err := db.QueryRow(query, mobile).Scan(&user.ID, &user.MobileNumber, &user.Name)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
