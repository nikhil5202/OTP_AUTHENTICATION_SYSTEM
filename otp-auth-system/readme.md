# OTP Authentication System

A Go-based OTP Authentication System that allows user registration, login, OTP verification, and device fingerprinting. It uses **Twilio** for sending OTPs to mobile numbers and **MySQL** as the database. The system also includes features for managing user sessions and ensuring secure access via OTP validation.

---

## Features

- **User Registration**: Register users with mobile number and name.
- **OTP Authentication**: Login users by verifying OTP sent to their mobile numbers.
- **Resend OTP**: Option to resend OTP if it expires or wasn't received.
- **User Details**: Fetch user details after successful login.
- **Device Fingerprinting**: Identifies and manages different devices from which users login.
- **Twilio Integration**: Secure OTP delivery via **Twilio SMS**.

---

## Technologies Used

- **Go (Golang)**: Backend development.
- **MySQL**: Database for user and OTP data.
- **Twilio API**: OTP sending via SMS.
- **Gorilla Mux**: HTTP routing.
- **Postman**: API testing.

---

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go** (Golang) v1.19 or higher
- **MySQL** database server running
- **Twilio Account** with API credentials for OTP sending

---

## Setup and Installation
1. Set up MySQL Database and Create a new MySQL database named otp_auth.
Run the SQL queries in db.sql to set up the necessary tables.

2. Configure environment variables:

Set up your Twilio credentials in the .env file or in your environment variables.
The variables you need to set are:
# TWILIO_ACCOUNT_SID=your_twilio_account_sid
# TWILIO_AUTH_TOKEN=your_twilio_auth_token
# TWILIO_PHONE_NUMBER=your_twilio_phone_number

3. Install Dependencies:

go mod tidy

4. Run the server:

go run main.go


5. Test the API:
Import the postman collection: OTP_Auth.postman_collection.json.
Use Postman testing tool to test the following endpoints:

POST /register: Register a new user with mobile number and name.
POST /login: Login by providing mobile number and OTP.
POST /resend-otp: Resend OTP to the user's mobile number.
GET /user/details: Fetch the logged-in user's details using user_id in the header.


Postman collection:
https://speeding-capsule-994459.postman.co/workspace/My-Workspace~2bff2343-1582-4759-8d47-f8f25b03d544/collection/39718258-24ff85c8-facb-4246-897d-238f7eba5bab?action=share&creator=39718258

## Important:

For Twilio Trial Account, real time otp is sent to only registered mobile number. OTP is also available in the response from the API.