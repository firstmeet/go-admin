package Upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
	"weserver/utils"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	file.Filename = "resource/images/" + strconv.Itoa(int(time.Now().UnixNano()))+utils.GetRandomString(4)+file.Filename[strings.LastIndex(file.Filename,"."):]
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"data": file.Filename,
	})
}
