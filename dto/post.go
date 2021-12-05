package dto

type CreatePostDto struct {
	UID     int    `json:"uid" gorm:"uid"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  int    `json:"status"`
}

type UpdatePostDto struct {
	ID      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  int    `json:"status"`
}
