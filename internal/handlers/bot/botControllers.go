package bot

import (
	"beak/internal/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// "beak/pkg/dataBase"
	"github.com/jinzhu/gorm"
)

const (
	IdentityJWTKey = "id"
)

type BotController struct {
	Database *gorm.DB
}

func (b *BotController) newBot(c *gin.Context) {
	var bot models.BotFront
	if err := c.ShouldBindJSON(&bot); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error new bot": err.Error()})
		return
	}
	body, err := json.Marshal(bot.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error body": err.Error()})
		return
	}
	botBd := models.Bot{
		ID:        bot.ID,
		CreatedAt: bot.CreatedAt,
		UpdatedAt: bot.UpdatedAt,
		DeletedAt: bot.DeletedAt,
		Id_user:   bot.Id_user,
		Token:     bot.Token,
		Body:      body,
	}
	if resErr := b.Database.Debug().Create(&botBd); resErr.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error create bot": resErr.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
func (b *BotController) getBots(c *gin.Context) {
	var bots models.Bots
	if err := c.ShouldBindJSON(&bots); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error get bots": err.Error()})
		return
	}
	botErr := b.Database.Debug().Where("id_user = ?", bots.Id_user).Find(&[]models.Bot{})
	fmt.Println(botErr.RowsAffected)
	if botErr.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "res": nil})
		return
	}
	// fmt.Println(botErr)
	c.JSON(http.StatusOK, gin.H{"success": true, "res": botErr})
}

func (b *BotController) getBot(c *gin.Context) {
	var bot models.BotFront
	var botBd models.Bot
	if err := c.ShouldBindJSON(&bot); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error get bot": err.Error()})
		return
	}
	botErr := b.Database.Debug().Where("id = ?", bot.ID).Find(&botBd)

	if botErr.Error != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This bot is not get"})
		return
	}
	fmt.Println(botBd)
	var body interface{}

	err := json.Unmarshal(botBd.Body, &body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This body is not get"})
		return
	}
	bot = models.BotFront{
		ID:        botBd.ID,
		CreatedAt: botBd.CreatedAt,
		UpdatedAt: botBd.UpdatedAt,
		DeletedAt: botBd.DeletedAt,
		Id_user:   botBd.Id_user,
		Token:     botBd.Token,
		Body:      body,
	}
	fmt.Println(bot)
	c.JSON(http.StatusOK, gin.H{"success": true, "res": bot})
}

func (b *BotController) saveBot(c *gin.Context) {
	var bot models.BotFront
	if err := c.ShouldBindJSON(&bot); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error save bot": err.Error()})
		return
	}
	body, err := json.Marshal(bot.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error body": err.Error()})
		return
	}
	botBd := models.Bot{
		ID:        bot.ID,
		CreatedAt: bot.CreatedAt,
		UpdatedAt: bot.UpdatedAt,
		DeletedAt: bot.DeletedAt,
		Id_user:   bot.Id_user,
		Token:     bot.Token,
		Body:      body,
	}
	fmt.Println(botBd.Body)
	botErr := b.Database.Model(&models.Bot{}).Debug().Where("id = ?", botBd.ID).Update(botBd)
	if botErr.Error != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "This tofen is not save"})
		return
	}
	fmt.Println(botErr.RowsAffected)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (b *BotController) getLogo(c *gin.Context) {

	// fmt.Println(reflect.TypeOf(f))

	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("./logo.png")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)

	c.JSON(http.StatusOK, gin.H{"success": true, "res": base64Encoding})
}

// if resErr := u.Database.Debug().Create(&user); resErr.Error != nil {
// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error create user": resErr.Error.Error()})
// 	return
// }
// c.JSON(http.StatusOK, gin.H{"success": true})
