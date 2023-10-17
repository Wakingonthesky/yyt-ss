package middleware

import (
	"fmt"
	"net/http"
	"time"
	"ytt-societyservice/controller/res"
	"ytt-societyservice/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
)

/*
Jwt 设计(仅用于access_token申请)：
检查是否携带refresh_token
若未携带，直接返回301状态码
若携带，检查refresh_token

	若无效或错误，返回302状态码
	若有效，进行下一步
*/
func JwtGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f models.Token
		if err := c.ShouldBind(&f); err != nil {
			fmt.Println("bind" + f.AccessToken)
			c.JSON(http.StatusOK, res.OKWithCode(res.BadRequest))
		} else {
			fmt.Println(f.AccessToken)
			if f.AccessToken == "" {
				c.JSON(http.StatusOK, res.OKWithCode(res.NotTakeToken))
				c.Abort()
				return
			} else {
				MySecret := []byte("mY9bY4xN4f")

				fmt.Println("feikong" + f.AccessToken)
				var claims models.MyClaims

				token, err := jwt.ParseWithClaims(f.AccessToken, &claims, func(t *jwt.Token) (interface{}, error) { return MySecret, nil })

				if err != nil {
					fmt.Println(token.Claims)
				}

				myClaims := token.Claims.(*models.MyClaims)

				if myClaims == nil {
					fmt.Println(claims)
					c.JSON(http.StatusOK, res.OKWithMessage("kong"))
					c.Abort()
					return
				}

				if time.Now().Unix() > myClaims.ExpiresAt.Unix() {
					fmt.Println(myClaims.ExpiresAt.Unix())
					c.JSON(http.StatusOK, res.OKWithCode(res.WrongToken))
					c.Abort()
					return
				} else {
					c.Next()
				}
			}
		}

	}
}

func JwtPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f models.Token
		if err := c.ShouldBindBodyWith(&f, binding.JSON); err != nil {
			fmt.Println("bind" + f.AccessToken)
			c.JSON(http.StatusOK, res.OKWithCode(res.BadRequest))
		} else {
			fmt.Println(f.AccessToken)
			if f.AccessToken == "" {
				c.JSON(http.StatusOK, res.OKWithCode(res.NotTakeToken))
				c.Abort()
				return
			} else {
				MySecret := []byte("mY9bY4xN4f")

				fmt.Println("feikong" + f.AccessToken)
				var claims models.MyClaims

				token, err := jwt.ParseWithClaims(f.AccessToken, &claims, func(t *jwt.Token) (interface{}, error) { return MySecret, nil })

				if err != nil {
					fmt.Println(token.Claims)
				}

				myClaims := token.Claims.(*models.MyClaims)

				if myClaims == nil {
					fmt.Println(claims)
					c.JSON(http.StatusOK, res.OKWithMessage("kong"))
					c.Abort()
					return
				}

				if time.Now().Unix() > myClaims.ExpiresAt.Unix() {
					fmt.Println(myClaims.ExpiresAt.Unix())
					c.JSON(http.StatusOK, res.OKWithCode(res.WrongToken))
					c.Abort()
					return
				} else {
					c.Next()
				}
			}
		}

	}
}
