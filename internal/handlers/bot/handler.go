package bot

import (
	"beak/internal/handlers"

	middleware "beak/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	saveBot      = "/save"
	saveBotToken = "/saveToken"
	getBots      = "/getBots"
	newBot       = "/newBot"
	getBot       = "/getBot"
	getLogo      = "/getLogo"
)

type botHandler struct {
	db *gorm.DB
}

func NewHandler(conn *gorm.DB) handlers.Handler {
	return &botHandler{db: conn}
}

func (bH *botHandler) Register(router *gin.RouterGroup) {
	bC := BotController{Database: bH.db}

	mwS := middleware.MiddlewareStorage{Database: bH.db}
	jwtMiddleware := mwS.JwtMiddleware(bH.db)
	router.POST(getLogo, bC.getLogo)
	router.Use(jwtMiddleware.MiddlewareFunc())
	{
		router.POST(newBot, bC.newBot)
		router.POST(getBots, bC.getBots)
		router.POST(getBot, bC.getBot)
		router.POST(saveBotToken, bC.saveBot)
		router.POST(saveBot, bC.saveBot)
		
		// router.GET(refrashToken, jwtMiddleware.RefreshHandler)
		// router.GET(readUser, bH.Read)
		// router.PUT(updateUser, bH.Update)

		// AdminGroup := router.Group(admin)
		// AdminGroup.Use(mwS.AdminMiddleware)
		// {
		// 	AdminGroup.DELETE(adminDelete, uC.Delete)
		// }
	}

}
