package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"guanxingtuan_bck/models"
	"guanxingtuan_bck/models/res"
	"time"
)

/*
Jwt 设计(仅用于access_token申请)：
检查是否携带refresh_token
若未携带，直接返回301状态码
若携带，检查refresh_token

	若无效，返回302状态码
	若有效，进行下一步
*/
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f models.Token
		if err := c.BindJSON(&f); err != nil {
			res.OKWithCode(500, c)
		} else {

			if f.Token == "" {
				res.OKWithCode(301, c)
			} else {
				claims := &models.MyClaims{}
				jwt.ParseWithClaims(f.Token, claims, func(t *jwt.Token) (interface{}, error) { return claims, nil })
				fmt.Print(claims)
				if time.Now().Unix() > claims.ExpiresAt.Unix() {
					res.OKWithCode(302, c)
					c.Abort()
					return
				} else {
					c.Next()
				}
			}
		}

	}
}
