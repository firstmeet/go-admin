package sqls

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var Db *gorm.DB
func RunSql(){
	fmt.Println("111111111111111111")
    config:=GetConfig()
    var err1 error
	Db, err1 = gorm.Open("mysql", config.db_user+":"+config.db_password+"@tcp("+config.db_host+")/"+config.db_name+"?charset=utf8&parseTime=True&loc=Local")
	Db.LogMode(true)
	if err1 !=nil {
		panic(err1)
	}
	//defer Db.Close()
}