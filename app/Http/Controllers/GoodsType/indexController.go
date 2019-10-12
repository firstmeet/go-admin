package GoodsType

import (
	"github.com/gin-gonic/gin"
	"time"
	"weserver/sqls"
	"weserver/utils"
)

type GoodsType struct {
	ID   int    `json:"id"`
	Name string `gorm:"type:varchar(30)" json:"name"`
	CreatedAt utils.JSONTime `json:"created_at"`
	UpdatedAt utils.JSONTime `json:"updated_at"`
}

func Index(c *gin.Context) {
	var goods_type []GoodsType
	sqls.Db.Find(&goods_type)
	c.JSON(200, gin.H{
		"data": goods_type,
	})
}
func Create(c *gin.Context) {
	var goods_type = GoodsType{
		Name: c.PostForm("name"),
		CreatedAt:utils.JSONTime{time.Now()},
	}
	sqls.Db.Create(&goods_type)
	c.JSON(200,gin.H{
		"data":goods_type,
	})
}
func Update(c *gin.Context)  {
	var goods_type GoodsType
	sqls.Db.Model(&goods_type).Where("id=?",c.Param("id")).Updates(map[string]interface{}{"name":c.PostForm("name"),"updated_at":utils.JSONTime{time.Now()}})
	c.JSON(200,gin.H{
		"data":"update success",
	})
}
func Delete(c *gin.Context){
	var goods_type GoodsType
	sqls.Db.Where("id=?",c.Param("id")).Delete(&goods_type)
	c.JSON(200,gin.H{
		"data":"delete success",
	})
}
