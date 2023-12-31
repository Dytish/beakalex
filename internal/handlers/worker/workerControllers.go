package worker

import (
	"encoding/base64"
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

type WorkerController struct {
	Database *gorm.DB
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (w *WorkerController) getWorker(c *gin.Context) {

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
