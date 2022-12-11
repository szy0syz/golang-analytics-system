package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type Event struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	EventType string
	Path      string
	Search    string
	Title     string
	URL       string
	CreatedAt int64 `gorm:"autoCreateTime"` // same as receivedAt
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

//func SaveEvent(db *gorom.DB, event Event) (Event, error) {
//	result := db.Clause()
//}

func main() {
	dbURL := `postgres://postgres:postgres@localhost:5432/postgres`

	ormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "kafka_",
		},
	}

	db, err := gorm.Open(postgres.Open(dbURL), ormConfig)
	if err != nil {
		panic(`Unable to connect to db`)
	}
	log.Println("=>Connected to successfully:", db)
	err = db.AutoMigrate(&Event{})
	if err != nil {
		log.Println("Error migrating schema:", err)
	}
}
