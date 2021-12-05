package utils

import (
	"encoding/json"

	"github.com/blog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSessionID(c *gin.Context) int {
	var id int
	session := sessions.Default(c)
	userInfo := session.Get("userinfo")
	userinfoStr, ok := userInfo.(string)
	if ok {
		var userinfoStruct []*model.User
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		id = userinfoStruct[0].Id
	}
	return id
}

func GetSessionRole(c *gin.Context) int16 {
	var role int16
	session := sessions.Default(c)
	userInfo := session.Get("userinfo")
	userinfoStr, ok := userInfo.(string)
	if ok {
		var userinfoStruct []*model.User
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		role = userinfoStruct[0].Role
	}
	return role
}
