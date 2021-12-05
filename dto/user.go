package dto

type CreateUserDto struct {
	Username string `json:"username" binding:"required" gorm:"username"`
	Password string `json:"password" binding:"required" gorm:"password"`
}
