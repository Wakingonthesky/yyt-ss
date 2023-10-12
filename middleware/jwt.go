package middleware

import (
	"time"
	"ytt-societyservice/controller/res"
	"ytt-societyservice/models"

	"github.com/gin-gonic/gin"
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
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f models.Token
		if err := c.BindJSON(&f); err != nil {
			res.OKWithCode(res.BadRequest)
		} else {

			if f.AccessToken == "" {
				res.OKWithCode(res.NotTakeToken)
				c.Abort()
				return
			} else {
				claims := &models.MyClaims{}
				jwt.ParseWithClaims(f.AccessToken, claims, func(t *jwt.Token) (interface{}, error) { return claims, nil })
				if claims.ID == "" {
					res.OKWithCode(res.WrongToken)
					c.Abort()
					return
				}
				if time.Now().Unix() > claims.ExpiresAt.Unix() {
					res.OKWithCode(res.WrongToken)
					c.Abort()
					return
				} else {
					c.Next()
				}
			}
		}

	}
}
