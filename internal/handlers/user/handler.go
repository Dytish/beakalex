package user

import (
	"beak/internal/handlers"

	middleware "beak/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	signUp       = "/signup"
	login        = "/login"
	refrashToken = "/refresh/token"
	readUser     = "/read"
	updateUser   = "/update"
)

type userHandler struct {
	db *gorm.DB
}

func NewHandler(conn *gorm.DB) handlers.Handler {
	return &userHandler{db: conn}
}

func (uH *userHandler) Register(router *gin.RouterGroup) {
	uC := UserController{Database: uH.db}
	mwS := middleware.MiddlewareStorage{Database: uH.db}
	jwtMiddleware := mwS.JwtMiddleware(uH.db)

	router.POST(signUp, uC.SignUp)
	router.POST(login, jwtMiddleware.LoginHandler)
	router.Use(jwtMiddleware.MiddlewareFunc())
	{
		router.GET(refrashToken, jwtMiddleware.RefreshHandler)
		router.GET(readUser, uC.Read)
		router.PUT(updateUser, uC.Update)

		// AdminGroup := router.Group(admin)
		// AdminGroup.Use(mwS.AdminMiddleware)
		// {
		// 	AdminGroup.DELETE(adminDelete, uC.Delete)
		// }
	}

}
