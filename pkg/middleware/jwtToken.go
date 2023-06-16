package middleware

import (
	"beak/internal/models"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	IdentityJWTKey = "id"
)

type JwtWrapper struct {
	SecretKey       string // Секретный ключ
	Issuer          string // Генерирующий токен
	ExpirationHours int64  // Срок действия
}

func (mwS *MiddlewareStorage) JwtMiddleware(conn *gorm.DB) *jwt.GinJWTMiddleware {
	m, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "diplom",
		Key:         []byte("superoleg"),
		Timeout:     time.Minute * 15,
		MaxRefresh:  time.Hour * 100,
		IdentityKey: IdentityJWTKey,
		RefreshResponse: func(c *gin.Context, code int, token string, t time.Time) {

			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"token":   token,
				"expire":  t.Format(time.RFC3339),
				"message": "refresh successfully",
			})
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					IdentityJWTKey: v.ID,
				}
			}
			return jwt.MapClaims{
				"error": true,
			}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			if v, ok := claims[IdentityJWTKey].(float64); ok {
				return &models.User{
					ID: uint(v),
				}
			}
			return &models.User{
				ID: 0,
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var credentials = struct {
				Login    string `form:"login" json:"login" binding:"required"`
				Password string `form:"password" json:"password" binding:"required"`
			}{}

			if err := c.ShouldBind(&credentials); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			var userModel models.User
			conn.Where(models.User{Login: credentials.Login}).First(&userModel)
			if userModel.ID == 0 {
				return "", jwt.ErrFailedAuthentication
			}
			err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(credentials.Password))
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}
			return &userModel, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenHeadName:     "Bearer ",
		TokenLookup:       "header: Authorization, query: token, cookie: jwt",
		TimeFunc:          time.Now,
		SendAuthorization: true,
	},
	)

	if err != nil {
		logrus.Errorf("Can't wake up JWT Middleware! Error: %s\n", err.Error())
		return nil
	}

	errInit := m.MiddlewareInit()
	if errInit != nil {
		logrus.Errorf("Can't init JWT Middleware! Error: %s\n", errInit.Error())
		return nil
	}

	return m
}
