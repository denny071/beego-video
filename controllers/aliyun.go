package controllers

import (
	"encoding/json"
	"fmt"
	"beego-video/models"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/astaxie/beego"
)

type AliyunController struct {
	beego.Controller
}

var (
	accessKeyId = "LTAI4G2YKJPj9T4kRK5bjwVU"
	accessKeySecret = "eTaASruNbTw8kCpImtlabFCKrRTirK"
)

type JSONS struct {
	RequestId string
	UploadAddress string
	UploadAuth string
	VideoId string
}

type PlayJSONS struct {
	PlayAuth string
}

type CallbackData struct {
	EventTime string
	EventType string
	VideoId string
	Status string
	Exteng string
	StreamInfos []CallbackStreamInfosData
}

type CallbackStreamInfosData struct {
	Status string
	Bitrate int
	Definition string
	Duration int
	Encrypt bool
	FileUrl string
	Format string
	Fps int
	Height int
	Size int
	Width int
	JobId string
}



func (c * AliyunController) Upload()  {
	title := c.GetString("title")
	desc := c.GetString("desc")
	fileName := c.GetString("fileName")
	coverUrl := c.GetString("coverUrl")
	tags := c.GetString("tags")

	if title == "" {
		c.Data["json"] = ReturnError(40001,"必须添加标题")
		c.ServeJSON()
	}
	if desc == "" {
		c.Data["json"] = ReturnError(40001,"必须添加描述")
		c.ServeJSON()
	}
	if fileName == "" {
		c.Data["json"] = ReturnError(40001,"必须添加文件名称")
		c.ServeJSON()
	}
	if coverUrl == "" {
		c.Data["json"] = ReturnError(40001,"必须添加文件地址")
		c.ServeJSON()
	}
	client, err := c.InitVodClient(accessKeyId,accessKeySecret)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	response, err := c.MyCreateUploadVideo(client, title, desc,fileName,coverUrl,tags)
	if err != nil {
		panic(err)
	}
	data := &JSONS{
		response.RequestId,
		response.UploadAddress,
		response.UploadAuth,
		response.VideoId,
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c * AliyunController) Refresh()  {
	videoId := c.GetString("videoId")
	if videoId == "" {
		c.Data["json"] = ReturnError(40001,"必须传入视频ID")
		c.ServeJSON()
	}
	client, err := c.InitVodClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	response, err := c.MyRefreshUploadVideo(client, videoId)
	if err != nil {
		panic(err)
	}

	data := &JSONS{
		response.RequestId,
		response.UploadAddress,
		response.UploadAuth,
		response.VideoId,
	}
	c.Data["json"] = data
	c.ServeJSON()
}


func (c * AliyunController) PlayAuth()  {
	videoId := c.GetString("videoId")
	if videoId == "" {
		c.Data["json"] = ReturnError(40001,"必须传入视频ID")
		c.ServeJSON()
	}
	client, err := c.InitVodClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	response, err := c.MyGetPlayAuth(client,videoId)
	if err != nil {
		panic(err)
	}
	data :=&PlayJSONS{
		response.PlayAuth,
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c * AliyunController) Callback()  {
	var ob CallbackData
	r := c.Ctx.Input.RequestBody
	json.Unmarshal(r, &ob)
	models.SaveAliyunVideo(ob.VideoId, string(r))
	c.Ctx.WriteString("success")
}

// 初始化
func (c *AliyunController) InitVodClient(accessKeyId string, accessKeySecret string)(client *vod.Client, err error)  {
	regionId := "cn-shanghai"
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}
	config := sdk.NewConfig()
	// 实拍是否自动重试
	config.AutoRetry = true
	// 最大重试次数
	config.MaxRetryTime = 3
	// 连接超时，单位：纳秒；默认为3秒
	config.Timeout = 3000000000

	return vod.NewClientWithOptions(regionId,config,credential)
}

// 上传视频
func (c *AliyunController) MyCreateUploadVideo(client *vod.Client, title string, desc string, fileName string,
	coverUrl string, tags string) (response *vod.CreateUploadVideoResponse,err error)  {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = title
	request.Description = desc
	request.FileName = fileName
	request.CoverURL = coverUrl
	request.Tags = tags
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

// 刷新上传视频
func (c *AliyunController) MyRefreshUploadVideo(client *vod.Client,videoId string) (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return  client.RefreshUploadVideo(request)
}

// 获得播放权限
func (c *AliyunController) MyGetPlayAuth(client *vod.Client, videoId string) (response *vod.GetVideoPlayAuthResponse, err error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetVideoPlayAuth(request)
}

