package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
	"time"
)


type VideoController struct {
	beego.Controller
}

// ChannelAdvert 频道页 - 获取顶部广告

func (c *VideoController) ChannelAdvert(){
	channelId, _ := c.GetInt("channelId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须制定频道")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelAdvert(channelId);

	if err == nil {
		c.Data["json"] = ReturnSuccess(0,"success",videos,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"请求数据失败，请稍后重试")
		c.ServeJSON()
	}

}

// ChannelHotList 频道页 - 获取正在热播视频

func (c *VideoController) ChannelHotList() {
	channelId, _ := c.GetInt("channelId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定ChannelVideo频道")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelHotList(channelId);

	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", videos, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}
}

// ChannelRecommendRegionList 频道页 - 获取频道地区推荐的视频

func (c *VideoController) ChannelRecommendRegionList(){
	channelId, _ := c.GetInt("channelId")
	regionId, _ := c.GetInt("regionId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定频道")
		c.ServeJSON()
	}
	if regionId == 0 {
		c.Data["json"] = ReturnError(4002,"必须指定地区")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId,regionId);

	if err == nil {
		c.Data["json"] = ReturnSuccess(0,"success",videos,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"没有相关内容")
		c.ServeJSON()
	}
}

// ChannelRecommendTypeList 频道页 - 获取频道类型推荐的视频
func (c *VideoController) ChannelRecommendTypeList(){
	channelId, _ := c.GetInt("channelId")
	typeId, _ := c.GetInt("typeId")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定频道")
		c.ServeJSON()
	}
	if typeId == 0 {
		c.Data["json"] = ReturnError(4002,"必须指定频道类型")
		c.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendTypeList(channelId,typeId);

	if err == nil {
		c.Data["json"] = ReturnSuccess(0,"success",videos,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"没有相关内容")
		c.ServeJSON()
	}
}

// ChannelVideo 根据传入参数获取视频列表
func (c *VideoController) ChannelVideo(){
	// 获取频道ID
	channelId, _ := c.GetInt("channelId")
	// 获取频道地区ID
	regionId,_ := c.GetInt("regionId")
	// 获取频道类型ID
	typeId,_ := c.GetInt("typeId")
	// 获取状态
	end := c.GetString("end")
	// 获取排序
	sort := c.GetString("sort")
	// 获取页码信息
	limit,_ := c.GetInt("limit")
	offset,_ := c.GetInt("offset")

	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定频道")
		c.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId,regionId,typeId,end,sort,offset,limit)

	if err == nil {
		c.Data["json"] = ReturnSuccess(0,"success",videos,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"没有相关内容")
		c.ServeJSON()
	}
}

func (c *VideoController) VideoInfo()  {
	id := c.Ctx.Input.Param(":id")
	c.Data["json"] = SuccessJsonStruct{Msg: "video id:" + id}
	c.ServeJSON()
}

func (c *VideoController) EpisodesList()  {
	id := c.Ctx.Input.Param(":id")
	c.Data["json"] = SuccessJsonStruct{Msg: "video id:" + id + " episodes list"}
	c.ServeJSON()
}


func (c *VideoController) Save()  {
	c.Data["json"] = SuccessJsonStruct{Msg: " save video "}
	c.ServeJSON()
}


func (c *VideoController) Search() {
	c.Data["json"] = SuccessJsonStruct{Msg: " Search"}
	c.ServeJSON()
}


//生成测试视频数据
func (c *VideoController) TestData() {
	var i = 1
	for {
		i++
		rand.Seed(time.Now().UnixNano())
		uidi := rand.Intn(10)
		uid := uidi + 10
		models.SaveVideo(strconv.Itoa(i)+"鸣人柯南一护路飞由诺阿斯塔"+strconv.Itoa(i), "蜡笔小新樱桃小丸子", 1, 2, 2, "/static/video/coverr-sparks-of-bonfire-1573980240958.mp4", uid, "")
		i++
		c.Data["json"] = SuccessJsonStruct{Msg: " create Test data，title："+"鸣人柯南一护路飞由诺阿斯塔"+strconv.Itoa(i)}
		c.ServeJSON()
	}
}

//导入ES脚本
func (c *VideoController) SendEs() {
	//_, data, _ := models.GetAllList()
	//for _, v := range data {
	//	body := map[string]interface{}{
	//		"id":                   v.Id,
	//		"title":                v.Title,
	//		"sub_title":            v.SubTitle,
	//		"add_time":             v.AddTime,
	//		"img":                  v.Img,
	//		"img1":                 v.Img1,
	//		"episodes_count":       v.EpisodesCount,
	//		"is_end":               v.IsEnd,
	//		"channel_id":           v.ChannelId,
	//		"status":               v.Status,
	//		"region_id":            v.RegionId,
	//		"type_id":              v.TypeId,
	//		"episodes_update_time": v.EpisodesUpdateTime,
	//		"comment":              v.Comment,
	//		"user_id":              v.UserId,
	//		"is_recommend":         v.IsRecommend,
	//	}
	//	es.EsAdd("fyouku_video", "video-"+strconv.Itoa(v.Id), body)
	//}
}



