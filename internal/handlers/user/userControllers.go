package user

import (
	"beak/internal/models"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	// "beak/pkg/dataBase"
	"github.com/jinzhu/gorm"
)

const (
	IdentityJWTKey = "id"
)

type UserController struct {
	Database *gorm.DB
}

func (u *UserController) SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//c.JSON(http.StatusBadRequest, err.Error()) todo Не пиши ошибки с ключем. Ошибка в идеале должна быть просто текстом. Либо ключ всегда должен быть один. Прикинь какого фронта будет разные ключи обрабатывать? Не прикольно совсем
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error registration": err.Error()}) // todo : используй abort только в мидлварях. Ниче критичного тут нет, но abort явно прерывает контекст выполнения и идеоматически его приянто использовать именно в них (выше пример правильного использования)
		return
	}
	// u.Database.Debug().Where()
	if userErr := u.Database.Debug().Where("email = ?", user.Email).First(&models.User{}); userErr.RowsAffected == 1 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This email is already in use"})
		return
	}

	if userErr := u.Database.Debug().Where("login = ?", user.Login).First(&models.User{}); userErr.RowsAffected == 1 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This login is already in use"})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error hashing password": err.Error()})
		return
	}
	user.Password = string(bytes)

	if resErr := u.Database.Debug().Create(&user); resErr.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error create user": resErr.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (u *UserController) Read(c *gin.Context) {
	var user models.User

	claims := jwt.ExtractClaims(c)
	UserID := claims[IdentityJWTKey]

	if result := u.Database.Debug().Where("id = ?", UserID).Find(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error - not found user": result.Error.Error()})
		return
	}

	userDTO, err := user.Convert()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error convert to DTO": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, &userDTO)
	}
}

func (u *UserController) Update(c *gin.Context) {
	var user models.User

	claims := jwt.ExtractClaims(c)
	UserID := claims[IdentityJWTKey].(float64)
	user.ID = uint(UserID)

	if shouldErr := c.ShouldBindJSON(&user); shouldErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"ShouldBindJSON error": shouldErr.Error()})
		return
	}

	result := u.Database.Debug().Where("id = ?", user.ID).Find(&models.User{})

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error - user not found": result.Error.Error()})
		return
	}
	if userErr := u.Database.Debug().Debug().Where("login = ? and id <> ?", user.Login, user.ID).First(&models.User{}); userErr.RowsAffected == 1 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This login is already in use"})
		return
	}
	if userErr := u.Database.Debug().Where("email = ? and id <> ?", user.Login, user.ID).First(&models.User{}); userErr.RowsAffected == 1 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This email is already in use"})
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error hashing password": err.Error()})
		return
	}
	user.Password = string(bytes)
	if updateResult := result.Updates(&user); updateResult.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error update data": updateResult.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (u *UserController) Delete(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user id provided"})
		return
	}
	_, err := strconv.Atoi(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error - userID id must be uint": err.Error()})
		return
	}
	result := u.Database.Debug().Where("id = ?", userId).Find(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error user not found": result.Error.Error()})
		return
	} else {
		if resErr := u.Database.Debug().Delete(&user); resErr.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error delete user": resErr.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
