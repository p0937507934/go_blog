package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/blog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	//get userinfo for session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//assert userinfo is string
	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []model.User
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if len(userinfoStruct) > 0 {
			c.Next()
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "Please Login First.",
		})
		c.Abort()
		return
	}
}

func AuthAdmin(c *gin.Context) {
	//get userinfo for session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//assert userinfo is string
	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []model.User
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if len(userinfoStruct) == 0 || userinfoStruct[0].Role == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Auth Failed, Not Admin.",
			})
			c.Abort()
			return
		}
		if len(userinfoStruct) > 0 && userinfoStruct[0].Role == 1 {
			c.Next()
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "Auth Failed, Not Admin.",
		})
		c.Abort()
		return
	}
}
