package handler

import (
	"check/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"os"
)

func DownloadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")

	// gen uuid code
	uid := uuid.NewV4().String()

	fmt.Println(uid)

	pwd, _ := os.Getwd()

	// save file by uuid as file name
	c.SaveUploadedFile(file, pwd + "/tmp/" + uid + ".docx")

	results := core.Execute("tmp/" + uid + ".docx")

	c.JSON(200, gin.H{
		"filename": file.Filename,
		"data": results,
	})
}
