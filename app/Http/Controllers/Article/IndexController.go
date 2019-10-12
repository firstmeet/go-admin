package Article

import (
	"github.com/gin-gonic/gin"
	"time"
	"weserver/sqls"
	"weserver/utils"
)

type Article struct {
	ID          int            `json:"id"`
	UserID      int            `json:"user_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Cover       string         `json:"cover"`
	Body        string         `gorm:"type:longtext" json:"body"`
	Status      int            `gorm:"type:tinyint" json:"status"`
	User        User           `json:"user"`
	CreatedAt   utils.JSONTime `json:"created_at"`
	UpdatedAt   utils.JSONTime `json:"updated_at"`
}
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Index(c *gin.Context) {
	var article []Article
	sqls.Db.Preload("User").Find(&article)
	c.JSON(200, gin.H{
		"data": article,
	})
}
func Update(c *gin.Context) {
	id := c.Param("id")
	var article Article
	sqls.Db.Model(&article).Where("id=?", id).Updates(map[string]interface{}{"title": c.PostForm("title"), "description": c.PostForm("description"), "body": c.PostForm("body"), "status": utils.StringToInt(c.PostForm("status")), "cover": c.PostForm("cover")})
	sqls.Db.Preload("User").Where("id=?", id).Find(&article)
	c.JSON(200, gin.H{
		"data": article,
	})
}
func Create(c *gin.Context) {
	article := Article{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Body:        c.PostForm("body"),
		Cover:       c.PostForm("cover"),
		UserID:      utils.StringToInt(c.PostForm("user_id")),
		CreatedAt:   utils.JSONTime{time.Now()},
		Status:      utils.StringToInt(c.PostForm("status")),
	}
	sqls.Db.Create(&article)
	c.JSON(200, gin.H{
		"msg": "create success",
	})
}
