package domain

import "time"

type User struct {
    ID           string    `json:"id"`          
    Name         string    `json:"name"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`             
    CreatedAt    time.Time `json:"created"`
    UpdatedAt    time.Time `json:"updated"`
}

type UserResponse struct {
    ID        string    `json:"id"`        
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}