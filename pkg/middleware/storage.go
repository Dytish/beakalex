package middleware

import (
	"github.com/jinzhu/gorm"
)

type MiddlewareStorage struct {
	Database *gorm.DB
}