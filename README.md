# MyRestoAuth – Restaurant Authentication Service

A secure **restaurant authentication backend** built with **Golang**, implementing email verification, password setup, JWT authentication, and refresh token sessions.

This service is designed with a **clean layered architecture**:

```
Handler → Service → Repository → Database
```

Features include:

* Restaurant signup
* Email verification
* Password setup
* Secure login
* JWT access tokens
* Refresh token sessions
* PostgreSQL database
* Docker support
* Production deployment

---

# Live Backend

Deployed on Render:

```
https://myrestoauth-1.onrender.com
```

---

# Authentication Flow

```
Signup
  ↓
Verification Email Sent
  ↓
User clicks verification link
  ↓
Set Password
  ↓
Login
  ↓
Access Token + Refresh Token issued
  ↓
Use access token for protected APIs
```

---

# API Endpoints

## 1. Signup

Registers a new restaurant and sends a verification email.

**Endpoint**

```
POST /api/users/signup
```

**URL**

```
https://myrestoauth-1.onrender.com/api/users/signup
```

**Request Body**

```json
{
  "restaurant_name": "MyRestoToday Restaurant",
  "username": "Haris",
  "email": "example@gmail.com"
}
```

**Response**

```json
{
  "id": "uuid",
  "restaurant_name": "MyRestoToday Restaurant",
  "username": "Haris",
  "email": "example@gmail.com",
  "email_verified": false
}
```

---

## 2. Verify Email

Verifies the user's email using the token sent via email.

**Endpoint**

```
GET /api/users/verify-email
```

**Example URL**

```
https://myrestoauth-1.onrender.com/api/users/verify-email?token=<verification_token>
```

**Query Parameter**

```
token=verification_token
```

**Response**

```json
{
  "message": "email verified successfully"
}
```

---

## 3. Set Password

Sets the password after email verification.

**Endpoint**

```
POST /api/users/set-password
```

**URL**

```
https://myrestoauth-1.onrender.com/api/users/set-password
```

**Request Body**

```json
{
  "token": "verification_token",
  "password": "strongpassword123"
}
```

**Response**

```json
{
  "message": "password set successfully"
}
```

---

## 4. Login

Authenticates the user and returns access and refresh tokens.

**Endpoint**

```
POST /api/users/login
```

**URL**

```
https://myrestoauth-1.onrender.com/api/users/login
```

**Request Body**

```json
{
  "email": "example@gmail.com",
  "password": "strongpassword123"
}
```

**Response**

```json
{
  "access_token": "jwt_access_token",
  "refresh_token": "jwt_refresh_token"
}
```

---

## 5. Refresh Token

Generates a new access token using a refresh token.

**Endpoint**

```
POST /api/users/refresh
```

**URL**

```
https://myrestoauth-1.onrender.com/api/users/refresh
```

**Request Body**

```json
{
  "refresh_token": "jwt_refresh_token"
}
```

**Response**

```json
{
  "access_token": "new_access_token",
  "expires_at": "timestamp"
}
```

---

# Tech Stack

* Golang
* Gin Web Framework
* PostgreSQL
* GORM ORM
* JWT Authentication
* SMTP Email Verification
* Docker
* Render Deployment

---

# Security Features

* Password hashing using bcrypt
* Email verification tokens with expiry
* Refresh token session storage
* JWT access tokens
* Secure authentication flow

---

---

# Running Locally

1. Clone repository

```
git clone <repo-url>
```

2. Configure environment variables

```
DB_URL=
PORT=
SMTP_HOST=
SMTP_PORT=
SMTP_EMAIL=
SMTP_APP_PASSWORD=
EMAIL_FROM_NAME=
JWT_ACCESS_SECRET=
JWT_REFRESH_SECRET=
JWT_ACCESS_EXPIRY_MINUTE=
JWT_REFRESH_EXPIRY_DAYS=
VERIFICATION_TOKEN_EXPIRY_HOURS=
```

3. Run the application

```
go run cmd/main.go
```

---

# Author

**Anandhu Rajan**

Backend Developer – Golang
