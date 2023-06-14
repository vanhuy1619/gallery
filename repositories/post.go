package repositories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadImages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Parse the multipart form
		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File upload large"})
			return
		}

		files := c.Request.MultipartForm.File["images"]
		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error": "File error"})
				return
			}
			defer src.Close()

			contentType := file.Header.Get("Content-Type")
			if !strings.HasPrefix(contentType, "image/") {
				c.JSON(http.StatusBadRequest, gin.H{"Error": "File not supported"})
				return
			}

			//create destination
			ext := filepath.Ext(file.Filename)
			filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
			destination := filepath.Join("./public/images/", filename)
			dst, err := os.Create(destination)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Upload file error"})
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, src)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error file"})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"success": "Upload file successful"})

	}
}
func uploadVideo(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
