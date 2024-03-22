package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm/common"
	"go-gin-gorm/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register 注册
func Register(c *gin.Context) {
	// 连接数据库
	db := common.GetDB()
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password

	// 数据校验
	var user model.User
	db.Where("phone_number = ?", phoneNumber).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "该手机号已注册",
		})
	}

	// 加密密码
	thicken, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	newUser := model.User{
		UserName:    userName,
		PhoneNumber: phoneNumber,
		Password:    string(thicken),
		Avatar:      "/images/default_avatar.png",
		Collects:    model.Array{},
		Following:   model.Array{},
		Fans:        0,
	}

	db.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// Login 登录
func Login(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	// 数据验证
	var user model.User
	db.Where("phone_number =?", phoneNumber).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

// GetInfo 登录后获取信息
func GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")
	// 返回用户信息
	//response.Success(c, gin.H{"id": user.(model.User).ID, "avatar": user.(model.User).Avatar}, "登录获取信息成功")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"id": user.(model.User).ID, "avatar": user.(model.User).Avatar},
		"msg":  "登录获取信息成功",
	})
}
