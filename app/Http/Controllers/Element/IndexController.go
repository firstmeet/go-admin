package Element

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"weserver/sqls"
	"weserver/utils"
)

type Element struct {
	ID        int            `gorm:"primary_key" json:"id"`
	Title     string         `gorm:"type:varchar(50)" json:"title"`
	Picture   string         `gorm:"type:varchar(100)" json:"picture"`
	Media     string         `gorm:"type:varchar(100)" json:"media"`
	Link      string         `gorm:"type:varchar(255)" json:"link"`
	Type      int            `gorm:"type:int(2)" json:"type"`
	BindType  int            `gorm:"type:int(2)" json:"bind_type"`
	BindID    int            `gorm:"type:int(11)" json:"bind_id"`
	Sort      int            `gorm:"type:int(11)" json:"sort"`
	Status    int            `gorm:"type:int(1)" json:"status"`
	CreatedAt utils.JSONTime `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt utils.JSONTime `gorm:"type:timestamp" json:"updated_at"`
}

func Index(c *gin.Context) {
	var elements []Element
	//var user_info UserInfo
	sqls.Db.Find(&elements)
	c.JSON(200, gin.H{
		"data": elements,
	})
}
func Info(c *gin.Context) {
	id, err := c.Get("id")
	if err == false {
		panic("未获取")
	}
	var element Element
	sqls.Db.Where("id=?", id).First(&element)
	c.JSON(200, gin.H{
		"data": element,
	})

}

/*
*创建element
 */
func Create(c *gin.Context) {
	var element = Element{
		Title:   c.PostForm("title"),
		Picture: c.PostForm("picture"),
		Media:   c.PostForm("media"),
		Link:c.PostForm("link"),
		Type:    utils.StringToInt(c.PostForm("type")),
		Sort:    utils.StringToInt(c.PostForm("Sort")),
		Status:  utils.StringToInt(c.PostForm("status")),
	}
	sqls.Db.Create(&element)
	c.JSON(200, gin.H{
		"data": element,
	})

}

/*
* 更新用户信息
 */
func Update(c *gin.Context) {
	var element Element
	sqls.Db.Model(&element).Where("id = ?", c.Param("id")).Updates(map[string]interface{}{"title": c.PostForm("title"), "picture": c.PostForm("picture"), "media": c.PostForm("media"),"link":c.PostForm("link"), "type": utils.StringToInt(c.PostForm("type")), "sort": utils.StringToInt(c.PostForm("sort")), "status": utils.StringToInt(c.PostForm("status"))})
	c.JSON(200, gin.H{
		"message": "update success",
	})

}
func Show(c *gin.Context) {
	name := c.PostForm("title")
	var element []Element
	sqls.Db.Where("title like ?", "%"+name+"%").Find(&element)
	c.JSON(200, gin.H{
		"data": element,
	})
}
