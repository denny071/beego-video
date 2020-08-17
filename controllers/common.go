package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/astaxie/beego"
)

// CommonController 控制器
type CommonController struct {
	beego.Controller
}

// JsonStruct 结构
type SuccessJsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}
// JsonStruct 结构
type ErrorJsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
}


// ReturnSuccess 返回正确json
func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *SuccessJsonStruct) {
	json = &SuccessJsonStruct{Code: code, Msg: msg, Items: items, Count: count}
	return
}

// ReturnError 返回错误json
func ReturnError(code int, msg interface{}) (json *ErrorJsonStruct) {
	json = &ErrorJsonStruct{Code: code, Msg: msg}
	return
}

// MD5V md5加密
func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

// 格式化时间
func DataFormat(times int64) string {
	video_time := time.Unix(times, 0)
	return video_time.Format("2006-01-02")
}