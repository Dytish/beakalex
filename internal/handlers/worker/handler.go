package worker

import (
	"beakalex/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	getWorker = "/listWorkers"
)

type workerHandler struct {
	db *gorm.DB
}

func NewHandler(conn *gorm.DB) handlers.Handler {
	return &workerHandler{db: conn}
}

func (wH *workerHandler) Register(router *gin.RouterGroup) {
	wC := WorkerController{Database: wH.db}
	// mwS := middleware.MiddlewareStorage{Database: wH.db}
	// jwtMiddleware := mwS.JwtMiddleware(wH.db)
	router.POST(getWorker, wC.getWorker)

}
