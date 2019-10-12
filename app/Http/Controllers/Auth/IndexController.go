package Auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	Admin2 "weserver/app/Http/Controllers/Admin"
	User2 "weserver/app/Http/Controllers/User"
)

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Admin struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var SecretKey = "12312312312313"

func Register(c *gin.Context) {
	id := User2.Create(c)
	if id != 0 {
		token, erro1 := CreateToken([]byte(SecretKey), "app", uint(id), false)
		if erro1 != nil {
			panic(erro1)
		}
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		c.JSON(403, gin.H{
			"error": "用户已存在",
		})
	}
}
func Login(c *gin.Context) {
	user := make(map[string]string)
	user["mobile"] = c.PostForm("mobile")
	user["password"] = c.PostForm("password")
	id,user_info := User2.FindByLoginParam(user)
	if id != 0 {
		token, erro1 := CreateToken([]byte(SecretKey), "app", uint(id), false)
		if erro1 != nil {
			panic(erro1)
		}
		c.JSON(200, gin.H{
			"token": token,
			"user_info":user_info,
		})
	} else {
		c.JSON(403, gin.H{
			"error": "unauthorized",
		})
	}
}

func WxAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authString := c.Request.Header.Get("Authorization")
		if authString == "" {
			c.JSON(403, gin.H{
				"error": "unauthorized",
			})
			c.AbortWithStatus(403)
		} else {
			kv := strings.Split(authString, " ")
			if len(kv) != 2 || kv[0] != "Bearer" {
				c.JSON(403, gin.H{
					"error": "unauthorized",
				})
			}

			tokenString := kv[1]

			// Parse token
			token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKey), nil
			})

			if err != nil {
				c.JSON(403, gin.H{
					"error": "unauthorized",
				})
			}

			if !token.Valid {
				c.JSON(401, gin.H{
					"error": "unauthorized",
				})
			}
			claims, ok := token.Claims.(*jwtCustomClaims)

			if !ok {
				c.JSON(403, gin.H{
					"error": "unauthorized",
				})
			}
			//将uid写入请求参数
			uid := claims.Uid
			is_admin := claims.Admin
			c.Set("uid", uid)
			c.Set("is_admin", is_admin)
		}

		c.Next()

	}
}
func CreateToken(SecretKey []byte, issuer string, Uid uint, admin bool) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		Uid,
		admin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}
func AdminLogin(c *gin.Context) {
	admin := make(map[string]string)
	admin["name"] = c.PostForm("name")
	admin["password"] = c.PostForm("password")
	id := Admin2.Login(admin)
	if id != 0 {
		token, erro1 := CreateToken([]byte(SecretKey), "app", uint(id), true)
		if erro1 != nil {
			panic(erro1)
		}
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		c.JSON(403, gin.H{
			"error": "unauthorized",
		})
	}
}
