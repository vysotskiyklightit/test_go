package users

import (
	"net/http"
	"test_go/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersView struct {
	ParentGroup *gin.RouterGroup
	DB          *gorm.DB
}

func (view *UsersView) RegisterView() {
	router := view.ParentGroup.Group("/users")

	router.GET("", func(c *gin.Context) {
		var users []User
		view.DB.Find(&users)
		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	router.POST("", func(c *gin.Context) {
		var user User

		common.LoadJSONToModel(c, &user)
		view.DB.Create(&user)

		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
	})

}
