package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"otp-auth-system/models"
	"otp-auth-system/utils"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func ValidateOTP(db *sql.DB, mobile_number string, otp_code string) (bool, error) {
    var otpCode models.OTPCode
    query := `SELECT id, user_id, otp_code, expires_at, is_used FROM otp_codes 
              WHERE otp_code = ? AND is_used = FALSE AND expires_at > NOW()`
    err := db.QueryRow(query, otp_code).Scan(&otpCode.ID, &otpCode.UserID, &otpCode.OTP, &otpCode.ExpiresAt, &otpCode.IsUsed)
    
    log.Printf("Validating OTP: mobile=%s, otp=%s, queryResult=%+v", mobile_number, otp_code, otpCode)

    if err != nil {
        log.Printf("Error fetching OTP: %v", err)
        return false, err
    }
    _, err = db.Exec("UPDATE otp_codes SET is_used = TRUE WHERE id = ?", otpCode.ID)
    if err != nil {
        log.Printf("Error updating OTP usage: %v", err)
        return false, err
    }

    return true, nil
}


func ResendOTP(db *sql.DB, mobile string) error {
    otp := utils.GenerateOTP()
    expiryTime := time.Now().Add(24 * time.Hour)

    _, err := db.Exec("INSERT INTO otp_codes (user_id, otp_code, expires_at) VALUES ((SELECT id FROM users WHERE mobile_number = ?), ?, ?)", mobile, otp, expiryTime)
    if err != nil {
        return err
    }
    err = SendOTPViaTwilio(mobile, otp)
    return err
}

func SendOTPViaTwilio(mobileNumber string, otp string) error {
    accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
    authToken := os.Getenv("TWILIO_AUTH_TOKEN")
    fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")

    log.Printf(accountSid,authToken,fromNumber)

    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })

    message := fmt.Sprintf("OTP_AUTH: Hii, this is your OTP code from Nikhil Srivastava: %s . Don't share it to anyone.", otp)
    params := &openapi.CreateMessageParams{}
    params.SetTo(mobileNumber)
    params.SetFrom(fromNumber)
    params.SetBody(message)

    _, err := client.Api.CreateMessage(params)
    return err
}

func GenerateAndStoreOTP(db *sql.DB, mobileNumber string) (string, error) {
    otp := utils.GenerateOTP() 
    expiryTime := time.Now().Add(24 * time.Hour)

    _, err := db.Exec("INSERT INTO otp_codes (user_id, otp_code, expires_at) VALUES ((SELECT id FROM users WHERE mobile_number = ?), ?, ?)", mobileNumber, otp, expiryTime)
    if err != nil {
        return "", err
    }

    return otp, nil
}

func StoreDeviceInfo(db *sql.DB, mobile string, deviceFingerprint string) error {
    var userID int
    err := db.QueryRow("SELECT id FROM users WHERE mobile_number = ?", mobile).Scan(&userID)
    if err != nil {
        log.Printf("Error retrieving user ID: %v", err)
        return err
    }

    var existingID int
    err = db.QueryRow("SELECT id FROM devices WHERE user_id = ? AND device_fingerprint = ?", userID, deviceFingerprint).Scan(&existingID)

    if err == sql.ErrNoRows {
        _, err := db.Exec("INSERT INTO devices (user_id, device_fingerprint) VALUES (?, ?)", userID, deviceFingerprint)
        return err
    } else if err != nil {
        return err
    }
    _, err = db.Exec("UPDATE devices SET last_login = CURRENT_TIMESTAMP WHERE id = ?", existingID)
    return err
}