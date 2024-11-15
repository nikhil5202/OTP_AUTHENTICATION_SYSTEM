// handlers/user.go
package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "otp-auth-system/models"
    "log"
)

func GetUserDetails(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("user_id") 
        if userID == "" {
            http.Error(w, "Missing or invalid user_id", http.StatusUnauthorized)
            return
        }

        var user models.User
        query := `SELECT id, name, mobile_number FROM users WHERE id = ?`
        err := db.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.MobileNumber)

        if err != nil {
            log.Printf("Error fetching user details: %v", err)
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(user)
    }
}
