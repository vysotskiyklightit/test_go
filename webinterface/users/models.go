package users

type User struct {
	Name    string `gorm:"unique" json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}
