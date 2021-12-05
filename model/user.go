package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"column:username;unique;not null"`
	Password  string    `json:"-" gorm:"password;not null"`
	Role      int16     `json:"role" gorm:"column:role;default:0"`
	Status    int16     `json:"status" gorm:"column:status;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at; type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at; type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"-" gorm:"deleted_at"`
}

func (u *User) TableName() string {
	return "user"
}
