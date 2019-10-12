package User

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"weserver/sqls"
	"weserver/utils"
)

type User struct {
	ID        int            `gorm:"primary_key" json:"id"`
	Nickname  string         `gorm:"type:varchar(50)" json:"nickname"`
	Mobile    string          `gorm:"type:varchar(11)" json:"mobile"`
	Sex       int8            `gorm:"type:tinyint(4)" json:"sex"`
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`
	Country   string         `gorm:"type:varchar(100)" json:"country"`
	Province  string         `gorm:"type:varchar(100)" json:"province"`
	City      string         `gorm:"type:varchar(40)" json:"city"`
	Password  string         `gorm:"type:varchar(255)" json:"password"`
	CreatedAt utils.JSONTime `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt utils.JSONTime `gorm:"type:timestamp" json:"updated_at"`
}

func Index(c *gin.Context) {
	var users []User
	//var user_info UserInfo
	sqls.Db.Select([]string{"id", "nickname", "avatar", "country", "province", "city", "created_at", "updated_at"}).Find(&users)
	//sqls.Db.Model(&users).Related(&user_info).Find(&user_info)
	c.JSON(200, gin.H{
		"data": users,
	})
}
func Info(c *gin.Context) {
	id, err := c.Get("uid")
	if err == false {
		panic("未获取到用户")
	}
	var user User
	sqls.Db.Where("id=?", id).First(&user)
	c.JSON(200, gin.H{
		"data": user,
	})

}

/*
*创建用户
 */
func Create(c *gin.Context) int {
	md5str1 := utils.MD5Password(c.PostForm("password")) //将[]byte转成16进制
	if CheckUserExits(c.PostForm("email")) {
		return 0
	}
	user := User{
		//Name:      c.PostForm("name"),
		Password:  md5str1,
		CreatedAt: utils.JSONTime{time.Now()},
		UpdatedAt: utils.JSONTime{time.Now()},
	}
	//fmt.Print(user)
	//fmt.Println(sqls.Db)
	sqls.Db.Create(&user)
	return user.ID
}

/*
* 更新用户信息
 */
func Update(c *gin.Context) {
	var user User
	//var user_info UserInfo
	sqls.Db.Where("id=?", c.Param("id")).First(&user)
	//user.Name = c.PostForm("name")
	user.Password = utils.MD5Password(c.PostForm("password"))
	user.UpdatedAt = utils.JSONTime{time.Now()}

	sqls.Db.Save(&user)

	//var user_info UserInfo
	var user2 User
	sqls.Db.Where("id=?", c.Param("id")).First(&user2)
	c.JSON(200, gin.H{
		"data": user2,
	})
	//sqls.Db.Where("user_id=?", c.Param("id")).First(&user_info)
	//sqls.Db.Model(&user_info).Updates(map[string]interface{}{"avatar": c.PostForm("avatar"), "country": c.PostForm("country"), "province": c.PostForm("province"), "city": c.PostForm("city"), "area": c.PostForm("area")})

}
func FindByLoginParam(param map[string]string)(int,User) {
	var user User
	password := utils.MD5Password(param["password"])
	sqls.Db.Where("mobile=? AND password=?", param["mobile"], password).First(&user)
	fmt.Print(user)
	return user.ID,user
}
func CheckUserExits(email string) bool {
	var user User
	if sqls.Db.Where("email=?", email).First(&user).RecordNotFound() {
		return false
	} else {
		return true
	}

}
func Show(c *gin.Context) {
	name := c.PostForm("name")
	var user []User
	sqls.Db.Where("name like ?", "%"+name+"%").Find(&user)
	c.JSON(200, gin.H{
		"data": user,
	})
}
