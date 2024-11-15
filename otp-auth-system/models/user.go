// models/user.go
package models

type User struct {
    ID           int    `json:"id"`
    MobileNumber string `json:"mobile_number"`
    Name         string `json:"name"`
}

type OTPCode struct {
    ID        int
    UserID    int
    OTP       string
    ExpiresAt string
    IsUsed    bool
}

type Device struct {
    ID               int
    UserID           int
    DeviceFingerprint string
    LastLogin        string
}
