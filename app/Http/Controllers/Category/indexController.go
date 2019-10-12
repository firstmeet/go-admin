package Category

import (
	"github.com/gin-gonic/gin"
	"time"
	"weserver/sqls"
	"weserver/utils"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `gorm:"type:varchar(30)" json:"name"`
	//Description string         `gorm:"type:varchar(255)" json:"description"`
	Cover       string         `gorm:"type:varchar(255)" json:"cover"`
	//Body        string         `gorm:"type:longtext" gorm:"type:longtext" json:"body"`
	Status    int            `gorm:"type:tinyint(1)" json:"status"`
	Sort      int            `gorm:"type:int(11)" json:"sort"`
	CreatedAt utils.JSONTime `json:"created_at"`
	UpdatedAt utils.JSONTime `json:"updated_at"`
}

func Index(c *gin.Context) {
	var category []Category
	sqls.Db.Find(&category)
	c.JSON(200, gin.H{
		"data": category,
	})
}
func Create(c *gin.Context) {
	var category = Category{
		Name: c.PostForm("name"),
		Sort:utils.StringToInt(c.PostForm("sort")),
		Cover: c.PostForm("cover"),
		//Description:c.PostForm("description"),
		//Body:   c.PostForm("body"),
		//Type:    utils.StringToInt(c.PostForm("type")),
		//Sort:    utils.StringToInt(c.PostForm("Sort")),
		Status: utils.StringToInt(c.PostForm("status")),
		CreatedAt:utils.JSONTime{time.Now()},
	}
	sqls.Db.Create(&category)
	c.JSON(200,gin.H{
		"data":category,
	})
}
func Update(c *gin.Context)  {
	var category Category
	sqls.Db.Model(&category).Where("id=?",c.Param("id")).Updates(map[string]interface{}{"name":c.PostForm("name"),"cover":c.PostForm("cover"),"sort":utils.StringToInt(c.PostForm("sort")),"status":utils.StringToInt(c.PostForm("status")),"updated_at":utils.JSONTime{time.Now()}})
	c.JSON(200,gin.H{
		"data":"update success",
	})
}
func Delete(c *gin.Context){
	var category Category
	sqls.Db.Where("id=?",c.Param("id")).Delete(&category)
	c.JSON(200,gin.H{
		"data":"delete success",
	})
}
