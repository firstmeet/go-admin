package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"weserver/sqls"
	"weserver/utils"
)

type Admin struct {
	ID        int            `gorm:"primary_key" json:"id"`
	Name      string         `gorm:"type:varchar(50)" json:"name"`
	Password  string         `gorm:"type:varchar(255)" json:"password"`
	CreatedAt utils.JSONTime `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt utils.JSONTime `gorm:"type:timestamp" json:"updated_at"`
	Avatar string `json:"avatar"`
}
func Login(param map[string]string)int{
	var admin Admin
	password := utils.MD5Password(param["password"])
	sqls.Db.Where("name=? AND password=?", param["name"], password).First(&admin)
	fmt.Print(admin)
	return admin.ID
}
func Info(c *gin.Context){
	id,err:=c.Get("uid")
	is_admin,err:=c.Get("is_admin")
	if is_admin ==false {
		c.JSON(403, gin.H{
			"error": "unauthorized",
		})
	}
	if err==false {
		panic("未获取到用户")
	}
	var admin Admin
	sqls.Db.Where("id=?",id).Select([]string{"id", "name","created_at", "updated_at","avatar"}).First(&admin)
	c.JSON(200, gin.H{
		"data": admin,
	})

}
func Update(c *gin.Context) {
	var admin Admin
	id,err:=c.Get("id")
	if err ==false {
		panic("unauthorized")
	}
	sqls.Db.Where("id=?", id).First(&admin)
	admin.Name = c.PostForm("name")
	admin.Password = utils.MD5Password(c.PostForm("password"))
	admin.UpdatedAt = utils.JSONTime{time.Now()}

	sqls.Db.Save(&admin)

	//var user_info UserInfo
	//sqls.Db.Where("user_id=?", c.Param("id")).First(&user_info)
	//sqls.Db.Model(&user_info).Updates(map[string]interface{}{"avatar": c.PostForm("avatar"), "country": c.PostForm("country"), "province": c.PostForm("province"), "city": c.PostForm("city"), "area": c.PostForm("area")})

}