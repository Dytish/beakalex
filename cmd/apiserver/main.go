package main

import (
	"beakalex/internal/handlers/worker"
	"beakalex/pkg/dataBase"
	"beakalex/pkg/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	port      = "8089" // todo: Не нужно указывать перед портом двоеточие. Лучше используй конкатенацию строк типа: HOST+ ":" +PORT
	apiWorker = "/api/worker"
)

func main() {
	db := dataBase.GetDB()
	mw := middleware.MiddlewareStorage{}
	router := gin.Default()

	router.Use(mw.CORSMiddleware)

	v1Group := router.Group(apiWorker)
	userController := worker.NewHandler(db)
	userController.Register(v1Group)

	fmt.Println("\n" + "Start server...")
	log.Fatalf("Can't start app: %v", router.Run(":"+port))
}
