package Goods

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"strings"
	"time"
	"weserver/app/Http/Controllers/Category"
	"weserver/app/Http/Controllers/GoodsType"
	"weserver/sqls"
	"weserver/utils"
)

type Goods struct {
	ID          int     `gorm:"primary_key" json:"id"`
	CategoryID  int     `gorm:"type:bigint(11)" json:"category_id"`
	Name        string  `gorm:"type:varchar(50)" json:"name"`
	Cover       string  `gorm:"type:json" json:"cover"`
	OtherSku    int     `gorm:"type:tinyint(1)" json:"other_sku"`
	GoodsTypeID int     `gorm:"type:int(11)" json:"goods_type_id"`
	GoodsType GoodsType.GoodsType `json:"goods_type"`
	Sku []Sku `json:"sku"`
	SalePrice   float64 `gorm:"type:decimal()" json:"sale_price"`
	OriginPrice float64 `gorm:"type:decimal()" json:"origin_price"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Recommend   int     `gorm:"type:tinyint(1)" json:"recommend"`
	Body        string  `gorm:"type:longtext" json:"body"`
	//Sort      int            `gorm:"type:int(11)" json:"sort"`
	Status    int               `gorm:"type:int(1)" json:"status"`
	CreatedAt utils.JSONTime    `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt utils.JSONTime    `gorm:"type:timestamp" json:"updated_at"`
	Category  Category.Category `json:"category"`
}
type Skus struct {
	Sku []struct{
		Name string `json:"name"`
		ID int `json:"id"`
		SpecId int `json:"spec_id"`

	}
	OriginPrice string `json:"origin_price"`
	SalePrice string `json:"sale_price"`
	Stock string `json:"stock"`
}
type Sku struct {
	ID          int     `gorm:"primary_key" json:"id"`
	GoodsID int `json:"goods_id"`
	SkuCode     string `gorm:"type:varchar(100)" json:"sku_code"`
	Spec        string `gorm:"type:json" json:"spec"`
	SpecText string `gorm:"type:varchar(255)" json:"spec_text"`
	OriginPrice float64 `gorm:"type:decimal" json:"origin_price"`
	SalePrice float64 `gorm:"type:decimal" json:"sale_price"`
	Stock int `gorm:"type:int(11)" json:"stock"`
	CreatedAt utils.JSONTime    `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt utils.JSONTime    `gorm:"type:timestamp" json:"updated_at"`
}
func Index(c *gin.Context) {
	var goods []Goods
	//var user_info UserInfo
	current_page := utils.StringToInt(c.Query("current_page"))
	page_size := utils.StringToInt(c.Query("page_size"))
	offset := page_size * (current_page - 1)
	sqls.Db.Set("gorm:auto_preload", true).Limit(page_size).Offset(offset).Find(&goods)
	c.JSON(200, gin.H{
		"data": goods,
	})
}

/*
*创建goods
 */
func Create(c *gin.Context) {
	var goods = Goods{
		Name:        c.PostForm("name"),
		Cover:       c.PostForm("cover"),
		CategoryID:  utils.StringToInt(c.PostForm("category_id")),
		OtherSku:    utils.StringToInt(c.PostForm("other_sku")),
		GoodsTypeID:utils.StringToInt(c.PostForm("goods_type_id")),
		OriginPrice: float64(utils.StringToInt(c.PostForm("origin_price"))),
		SalePrice:   float64(utils.StringToInt(c.PostForm("sale_price"))),
		Recommend:   utils.StringToInt(c.PostForm("recommend")),
		Description: c.PostForm("description"),
		Body:        c.PostForm("body"),
		//Sort:    utils.StringToInt(c.PostForm("Sort")),
		Status:    utils.StringToInt(c.PostForm("status")),
		CreatedAt: utils.JSONTime{time.Now()},
	}
	sku_list:=c.PostForm("sku")
	if sku_list !="" {
		var sku []Skus
		json.Unmarshal([]byte(sku_list),&sku)
		var skus Sku
		sqls.Db.Where("goods_id=?",c.Param("id")).Delete(&skus)
		k:=0
		for _,v:=range sku{
			ids:=[]string{}
			fmt.Print(len(v.Sku))
			name:=[]string{}
			for _,v1:=range v.Sku{
				ids=append(ids,strconv.Itoa(v1.SpecId)+":"+strconv.Itoa(v1.ID))
				name=append(name,v1.Name)
			}
			if k!=0{
				k=k+1
			}

			name_string:=strings.Join(name,",")
			skus.Spec=string(utils.JSON(ids)[0:])
			skus.ID=k
			skus.Stock=utils.StringToInt(v.Stock)
			skus.SpecText=name_string
			skus.OriginPrice=float64(utils.StringToInt(v.OriginPrice))
			skus.SalePrice=float64(utils.StringToInt(v.SalePrice))
			skus.GoodsID=utils.StringToInt(c.Param("id"))
			skus.SkuCode=strconv.Itoa(int(time.Now().UnixNano()))
			skus.CreatedAt=utils.JSONTime{time.Now()}
			sqls.Db.Create(&skus)
		}
	}
	//var category Category.Category
	sqls.Db.Create(&goods).Related(&goods.Category)
	c.JSON(200, gin.H{
		"data": goods,
	})

}

/*
* 更新用户信息
 */
func Update(c *gin.Context) {
	var goods Goods
	sqls.Db.Model(&goods).Where("id = ?", c.Param("id")).Updates(map[string]interface{}{"name": c.PostForm("name"), "cover": c.PostForm("cover"), "description": c.PostForm("description"), "category_id": utils.StringToInt(c.PostForm("category_id")), "sale_price": float64(utils.StringToInt(c.PostForm("sale_price"))), "origin_price": float64(utils.StringToInt(c.PostForm("origin_price"))), "body": c.PostForm("body"), "recommend": utils.StringToInt(c.PostForm("recommend")), "other_sku": utils.StringToInt(c.PostForm("other_sku")),"goods_type_id":utils.StringToInt(c.PostForm("goods_type_id")), "status": utils.StringToInt(c.PostForm("status"))})
	//var goods2 Goods
	sku_list:=c.PostForm("sku")
	if sku_list !="" {
		var sku []Skus
		json.Unmarshal([]byte(sku_list),&sku)
		var skus Sku
		sqls.Db.Where("goods_id=?",c.Param("id")).Delete(&skus)
		k:=0
		for _,v:=range sku{
			ids:=[]string{}
			fmt.Print(len(v.Sku))
			name:=[]string{}
			for _,v1:=range v.Sku{
				ids=append(ids,strconv.Itoa(v1.SpecId)+":"+strconv.Itoa(v1.ID))
				name=append(name,v1.Name)
			}
			if k!=0{
				k=k+1
			}

			name_string:=strings.Join(name,",")
			skus.Spec=string(utils.JSON(ids)[0:])
			skus.ID=k
			skus.Stock=utils.StringToInt(v.Stock)
			skus.SpecText=name_string
			skus.OriginPrice=float64(utils.StringToInt(v.OriginPrice))
			skus.SalePrice=float64(utils.StringToInt(v.SalePrice))
			skus.GoodsID=utils.StringToInt(c.Param("id"))
			skus.SkuCode=strconv.Itoa(int(time.Now().UnixNano()))
			skus.CreatedAt=utils.JSONTime{time.Now()}
			sqls.Db.Create(&skus)
		}
	}


	sqls.Db.Set("gorm:auto_preload", true).Where("id=?", c.Param("id")).First(&goods)
	c.JSON(200, gin.H{
		"message": "update success",
		"data":    goods,
	})

}
func Show(c *gin.Context) {
	name := c.PostForm("title")
	var element []Goods
	sqls.Db.Where("title like ?", "%"+name+"%").Find(&element)
	c.JSON(200, gin.H{
		"data": element,
	})
}
func Delete(c *gin.Context) {
	var goods Goods
	sqls.Db.Where("id=?", c.Param("id")).Delete(&goods)
	c.JSON(200, gin.H{
		"data": "delete success",
	})
}
