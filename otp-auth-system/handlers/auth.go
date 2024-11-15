// handlers/auth.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"otp-auth-system/models"
	"otp-auth-system/services"
    "otp-auth-system/utils"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        ipAddress := utils.GetIP(r)

        // Optionally, store the IP address or include it in the response
        log.Printf("User registered from IP: %s", ipAddress)

        if err := services.CreateUser(db, user); err != nil {
            log.Printf("Error creating user: %v", err)
            http.Error(w, "Error creating user", http.StatusInternalServerError)
            return
        }

        otp, err := services.GenerateAndStoreOTP(db, user.MobileNumber)
        if err != nil {
            log.Printf("Error generating OTP: %v", err)
            http.Error(w, "Error generating OTP", http.StatusInternalServerError)
            return
        }

        if err := services.SendOTPViaTwilio(user.MobileNumber, otp); err != nil {
            log.Printf("Error sending OTP: %v", err)
            http.Error(w, "Error sending OTP", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "User registered successfully. OTP sent to mobile number",
            "user":    user,
            "otp":     otp,
            "ip":      ipAddress,
        })
    }
}


func LoginUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var requestData map[string]string
        if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
            log.Printf("Error decoding request: %v", err)  
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        log.Printf("Received request: %+v", requestData)  

        mobile_number := requestData["mobile_number"]
        otp_code := requestData["otp_code"]

        ipAddress := utils.GetIP(r)
        log.Printf("User logging in from IP: %s", ipAddress)

        deviceFingerprint := utils.GenerateDeviceFingerprint(r)
        log.Printf("Device Fingerprint: %s", deviceFingerprint)
        
        isValid, err := services.ValidateOTP(db, mobile_number, otp_code)
        if err != nil || !isValid {
            http.Error(w, "Invalid or expired OTP", http.StatusUnauthorized)
            return
        }

        user, err := services.GetUserByMobile(db, mobile_number)
        if err != nil {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }

        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Login successful",
            "user":    user,
            "device_fingerprint": deviceFingerprint,
            "ip":      ipAddress,
        })
    }
}

func ResendOTP(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var requestData map[string]string
        if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
            log.Printf("Error decoding request: %v", err)
            http.Error(w, "Invalid input", http.StatusBadRequest)
            return
        }

        mobile := requestData["mobile"]

        if mobile == "" {
            http.Error(w, "Mobile number is required", http.StatusBadRequest)
            return
        }

        err := services.ResendOTP(db, mobile)
        if err != nil {
            log.Printf("Error resending OTP: %v", err)
            http.Error(w, "Error resending OTP", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "OTP sent successfully"})
    }
}
