package models

type Uuid int

type Article struct {
	Id          Uuid   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Category struct {
	Id          Uuid   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	Id    Uuid   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Status struct {
	Error      string      `json:"error,omitempty"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data,omitempty"`
}
