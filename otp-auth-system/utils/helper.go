// utils/helper.go
package utils

import (
    "database/sql"
    "fmt"
    "math/rand"
    "time"
    "net"
    "strings"
    "net/http"
    "crypto/sha256"
    "encoding/hex"

    _ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
    dsn := "root:12345@tcp(127.0.0.1:3306)/otp_auth"

    return sql.Open("mysql", dsn)
}

func GenerateOTP() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("%06d", rand.Intn(1000000))
}


func GetIP(r *http.Request) string {
    ip := r.Header.Get("X-Forwarded-For")
    if ip == "" {
        ip, _, _ = net.SplitHostPort(r.RemoteAddr) 
    }
    ipList := strings.Split(ip, ",")
    return ipList[0]
}

func GenerateDeviceFingerprint(r *http.Request) string {
    ip := r.Header.Get("X-Forwarded-For")
    if ip == "" {
        ip = strings.Split(r.RemoteAddr, ":")[0]
    }
    
    userAgent := r.UserAgent()
    hash := sha256.New()
    hash.Write([]byte(ip + userAgent))
    return hex.EncodeToString(hash.Sum(nil))
}