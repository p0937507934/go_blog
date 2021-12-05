package dto

type CreateCommentDto struct {
	UID     int    `json:"uid" gorm:"uid"`
	Content string `json:"content" binding:"required"`
	PID     int    `json:"pid" binding:"required" gorm:"pid"`
	Status  int    `json:"status"`
}

type UpdateCommentDto struct {
	ID      int    `json:"id"`
	Content string `json:"content" binding:"required"`
	Status  int    `json:"status"`
}
