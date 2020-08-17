package controllers

import (
	"fyoukuApi/models"
	"regexp"

	"github.com/astaxie/beego"
)


type UserController struct {
	beego.Controller
}


func (c *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = c.GetString("mobile")
	password = c.GetString("password")

	if mobile == "" {
		c.Data["json"] = ReturnError(4001, "手机号不能为空")
		c.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		c.Data["json"] = ReturnError(4002, "手机格式不正确")
		c.ServeJSON()
	}

	if password == "" {
		c.Data["json"] = ReturnError(4003, "密码不能为空")
		c.ServeJSON()
	}

	// 判断手机号是否已经注册
	status := models.IsUserMobile(mobile)
	if status {
		c.Data["json"] = ReturnError(4004, "此手机已经注册")
		c.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err == nil {
			c.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			c.ServeJSON()
		} else {
			c.Data["json"] = ReturnError(5000, err)
			c.ServeJSON()
		}
	}
}

// LoginDo 用户登录
func (c *UserController) LoginDo() {
	mobile := c.GetString("mobile")
	password := c.GetString("password")

	if mobile == "" {
		c.Data["json"] = ReturnError(4001, "手机号不能为空")
		c.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|6|7)[0-9]\d{8}$`, mobile)
	if !isorno {
		c.Data["json"] = ReturnError(4002, "手机号格式不正确")
		c.ServeJSON()
	}
	if password == "" {
		c.Data["json"] = ReturnError(4003, "密码不能为空")
		c.ServeJSON()
	}
	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		c.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		c.ServeJSON()
	}
}


func (c *UserController) Video(){
	c.Data["json"] = SuccessJsonStruct{Msg: "video"}
	c.ServeJSON()
}
