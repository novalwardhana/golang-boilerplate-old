package model

const PasswordHashCost int = 10

type Params struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
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

type UserList struct {
	Total    int         `json:"total"`
	PerPage  int         `json:"per_page"`
	Page     int         `json:"page"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
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
