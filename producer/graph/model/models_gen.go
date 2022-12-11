// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Event struct {
	ID        string  `json:"id"`
	EventType *string `json:"eventType"`
	Path      *string `json:"path"`
	Search    *string `json:"search"`
	Title     *string `json:"title"`
	URL       *string `json:"url"`
	UserID    *string `json:"userId"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type PingResponse struct {
	Message string `json:"message"`
}

type RegisterKafkaEventInput struct {
	EventType string `json:"eventType"`
	UserID    string `json:"userId"`
	Path      string `json:"path"`
	Search    string `json:"search"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
