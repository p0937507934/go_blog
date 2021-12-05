package model

import "time"

type Comment struct {
	Id        int       `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Content   string    `json:"content" gorm:"content;not null"`
	PID       int       `json:"-" gorm:"pid;not null"`
	UID       int       `json:"-" gorm:"uid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UID"`
	Post      Post      `json:"post" gorm:"foreignKey:PID"`
	Status    int       `json:"status" gorm:"status;default:0;comment:'0 is normal, 1 is delete'"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at; type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"created_at; type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"-" gorm:"deleted_at"`
}

func (c *Comment) TableName() string {
	return "comment"
}
