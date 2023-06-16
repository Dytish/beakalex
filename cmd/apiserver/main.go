package main

import (
	"beak/internal/handlers/bot"
	"beak/internal/handlers/user"
	"beak/pkg/dataBase"
	"beak/pkg/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	port    = "8086" // todo: Не нужно указывать перед портом двоеточие. Лучше используй конкатенацию строк типа: HOST+ ":" +PORT
	apiUser = "/api/user"
	apiBot  = "/api/constructor"
)

func main() {
	db := dataBase.GetDB()
	mw := middleware.MiddlewareStorage{}
	router := gin.Default()

	router.Use(mw.CORSMiddleware)

	v1Group := router.Group(apiUser)
	userController := user.NewHandler(db)
	userController.Register(v1Group)

	v2Group := router.Group(apiBot)
	botController := bot.NewHandler(db)
	botController.Register(v2Group)

	fmt.Println("\n" + "Start server...")
	log.Fatalf("Can't start app: %v", router.Run(":"+port))
}
