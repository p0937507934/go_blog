package model

import "time"

type Post struct {
	Id        int       `json:"id" gorm:"id;primaryKey;autoIncrement"`
	UID       int       `json:"-" gorm:"uid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UID"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PID"`
	Title     string    `json:"title" gorm:"title;not null"`
	Content   string    `json:"content" gorm:"content;not null"`
	Status    int       `json:"status" gorm:"column:status;default:0;comment:'0 is normal, 1 is lock, 2 is delete'"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at; type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at; type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"-" gorm:"deleted_at"`
}

func (p *Post) TableName() string {
	return "post"
}
