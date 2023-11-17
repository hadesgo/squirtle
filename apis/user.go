package apis

import (
	"net/http"
	model "squirtle/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	var user model.User
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func Add(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    id,
		"message": "添加成功",
	})
}

func Update(c *gin.Context) {
	var user model.User
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(uint(id))
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    result,
		"message": "修改成功",
	})
}

func Delete(c *gin.Context) {
	var user model.User
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	result, err := user.Delete(uint(id))
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    result,
		"message": "删除成功",
	})
}
