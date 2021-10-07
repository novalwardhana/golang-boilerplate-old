package model

import "mime/multipart"

type UploadFile struct {
	File    *multipart.FileHeader
	FileExt string
}

type Result struct {
	Data  interface{}
	Error error
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type File struct {
	Directory string `json:"directory"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const PasswordHashCost int = 10
