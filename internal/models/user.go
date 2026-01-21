package models

import "time"

type User struct {
    ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
    Name      string    `bson:"name" json:"name"`
    Email     string    `bson:"email" json:"email"`
    Password  string    `bson:"password" json:"-"` // Don't expose password in JSON
    Role      string    `bson:"role" json:"role"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type SignupRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

type SigninRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type AuthResponse struct {
    Message string `json:"message"`
    Token   string `json:"token,omitempty"`
    User    *User  `json:"user,omitempty"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

