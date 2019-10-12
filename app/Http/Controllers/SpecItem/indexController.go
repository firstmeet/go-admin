package SpecItem

import (
	"github.com/gin-gonic/gin"
	"weserver/sqls"
	"weserver/utils"
)

type SpecItem struct {
	ID   int    `json:"id"`
	SpecID int `json:"spec_id"`
	Item string `gorm:"type:varchar(30)" json:"item"`
	CreatedAt utils.JSONTime `json:"created_at"`
	UpdatedAt utils.JSONTime `json:"updated_at"`
}

func Index(c *gin.Context) {
	var spec_item []SpecItem
	sqls.Db.Where("spec_id=?").Find(&spec_item)
	c.JSON(200, gin.H{
		"data": spec_item,
	})
}
//func Create(c *gin.Context) {
//	var spec_item = SpecItem{
//		Item: c.PostForm("name"),
//		CreatedAt:utils.JSONTime{time.Now()},
//	}
//	sqls.Db.Create(&goods_type)
//	c.JSON(200,gin.H{
//		"data":goods_type,
//	})
//}
//func Update(c *gin.Context)  {
//	var goods_type GoodsType
//	sqls.Db.Model(&goods_type).Where("id=?",c.Param("id")).Updates(map[string]interface{}{"name":c.PostForm("name"),"updated_at":utils.JSONTime{time.Now()}})
//	c.JSON(200,gin.H{
//		"data":"update success",
//	})
//}
//func Delete(c *gin.Context){
//	var goods_type GoodsType
//	sqls.Db.Where("id=?",c.Param("id")).Delete(&goods_type)
//	c.JSON(200,gin.H{
//		"data":"delete success",
//	})
//}
