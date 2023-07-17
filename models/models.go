package models

import "time"

type User struct {
	Email     string     `bson:"email,omitempty"`
	Password  string     `bson:"password,omitempty"`
	Platform  []string   `bson:"platform,omitempty"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt time.Time  `bson:"updatedAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
}
