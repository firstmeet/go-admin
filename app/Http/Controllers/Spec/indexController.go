package Spec

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"weserver/app/Http/Controllers/GoodsType"
	"weserver/sqls"
	"weserver/utils"
)

type Spec struct {
	ID   int    `json:"id"`
	Name string `gorm:"type:varchar(30)" json:"name"`
	GoodsTypeID int `gorm:"type:int(11)" json:"goods_type_id"`
	InputList string `gorm:"type:string" json:"input_list"`
	Sort int `gorm:"type:int(11)" json:"sort"`
	GoodsType GoodsType.GoodsType `json:"goods_type"`
	Status int `gorm:"type:int(1)" json:"status"`
	SpecItem []SpecItem `json:"spec_item"`
	CreatedAt utils.JSONTime `json:"created_at"`
	UpdatedAt utils.JSONTime `json:"updated_at"`
}
type SpecItem struct {
	ID   int    `gorm:"primary_key" json:"id"`
	SpecID int `json:"spec_id"`
	Item string `gorm:"type:varchar(30)" json:"item"`
	CreatedAt utils.JSONTime `json:"created_at"`
	UpdatedAt utils.JSONTime `json:"updated_at"`
}

func Index(c *gin.Context) {
	var spec []Spec

	is_pag:=c.Query("pag")
	goods_type_id:=c.Query("goods_type_id")
	if is_pag=="1"{
		current_page:=utils.StringToInt(c.Query("current_page"))
		page_size:=utils.StringToInt(c.Query("page_size"))
		offset:=page_size*(current_page-1)
		sqls.Db.Set("gorm:auto_preload", true).Where("goods_type_id=?",goods_type_id).Limit(page_size).Offset(offset).Find(&spec)
	}else{
		sqls.Db.Set("gorm:auto_preload", true).Where("goods_type_id=?",goods_type_id).Find(&spec)
	}

	c.JSON(200, gin.H{
		"data": spec,
	})
}
func Create(c *gin.Context) {
	var spec = Spec{
		Name: c.PostForm("name"),
		GoodsTypeID:utils.StringToInt(c.PostForm("goods_type_id")),
		InputList:c.PostForm("input_list"),
		Sort:utils.StringToInt(c.PostForm("sort")),
		Status:utils.StringToInt(c.PostForm("status")),
		CreatedAt:utils.JSONTime{time.Now()},
	}
	sqls.Db.Create(&spec).Related(&spec.GoodsType)
	var spec_item SpecItem
	list:=strings.Split(c.PostForm("input_list"),"\n")
	key:=0
	for _,v:= range list{
		if key!=0{
			spec_item.ID=key+1
		}
		spec_item.Item=v
		spec_item.SpecID=spec.ID
		spec_item.CreatedAt=utils.JSONTime{time.Now()}
		sqls.Db.Create(&spec_item)
		key=spec_item.ID
	}
	c.JSON(200,gin.H{
		"data":spec,
	})
}
func Update(c *gin.Context)  {
	var spec Spec
	sqls.Db.Model(&spec).Where("id=?",c.Param("id")).Updates(map[string]interface{}{"name":c.PostForm("name"),"goods_type_id":utils.StringToInt(c.PostForm("goods_type_id")),"input_list":c.PostForm("input_list"),"sort":utils.StringToInt(c.PostForm("sort")),"status":utils.StringToInt(c.PostForm("status")),"updated_at":utils.JSONTime{time.Now()}})
	var spec_item SpecItem
	sqls.Db.Where("spec_id=?",c.Param("id")).Delete(&spec_item)
	list:=strings.Split(c.PostForm("input_list"),",")
	key:=0
	for _,v:= range list{
		if key!=0{
			spec_item.ID=key+1
		}
		spec_item.Item=v
		spec_item.SpecID=utils.StringToInt(c.Param("id"))
		spec_item.CreatedAt=utils.JSONTime{time.Now()}
		sqls.Db.Create(&spec_item)
		key=spec_item.ID
	}
	sqls.Db.Set("gorm:auto_preload", true).Where("id=?",c.Param("id")).First(&spec)
	c.JSON(200,gin.H{
		"data":spec,
	})
}
func Delete(c *gin.Context){
	var spec Spec
	sqls.Db.Where("id=?",c.Param("id")).Delete(&spec)
	var spec_item SpecItem
	sqls.Db.Where("spec_id=?",c.Param("id")).Delete(&spec_item)
	c.JSON(200,gin.H{
		"data":"delete success",
	})
}
func FindByType(c *gin.Context){
	goods_type_id:=c.Query("goods_type_id")
	var spec []Spec
	sqls.Db.Set("gorm:auto_preload", true).Where("goods_type_id=?",goods_type_id).Find(&spec)
	c.JSON(200,gin.H{
		"data":spec,
	})
}
