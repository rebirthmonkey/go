package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	Username  string
	Firstname string
	Lastname  string
}

var identityKey = "id"

func main() {
	ginEngine := gin.New()

	basicStrategy := NewBasicStrategy(func(username string, password string) bool {
		if username == "admin" && password == "admin" {
			return true
		}
		return false
	})

	basicGroup := ginEngine.Group("/ping/basic")
	basicGroup.Use(basicStrategy.AuthFunc())
	{
		basicGroup.GET("/", func(c *gin.Context) {
			usernameKey := c.MustGet(UsernameKey).(string)
			c.JSON(http.StatusOK, gin.H{
				"UsernameKey": usernameKey,
				"message":     "pong",
			})
		})
	}

	ginJWT, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			if username == "admin" && password == "admin" {
				return &User{
					Username:  username,
					Lastname:  "Sun",
					Firstname: "Wukong",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username: claims[identityKey].(string),
			}
		},
		IdentityKey: identityKey,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.Username == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
	})
	jwtStrategy := NewJWTStrategy(*ginJWT)

	ginEngine.POST("/login/jwt", jwtStrategy.LoginHandler)

	jwtGroup := ginEngine.Group("/ping/jwt")
	jwtGroup.Use(jwtStrategy.AuthFunc())
	{
		jwtGroup.GET("/", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			log.Println("the claims is:", claims)
			user, _ := c.Get(identityKey)
			c.JSON(200, gin.H{
				"userID":   claims[identityKey],
				"userName": user.(*User).Username,
				"message":  "pong",
			})
		})
	}

	ginEngine.Run(":8080")
}
