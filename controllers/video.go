package controllers

import (
	"encoding/json"
	"fyoukuApi/models"
	"fyoukuApi/services/es"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)


type VideoController struct {
	beego.Controller
}

// ChannelAdvert 频道页-获取顶部广告
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

// ChannelHotList 频道页-获取正在热播视频
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

// ChannelRecommendRegionList 频道页-获取频道地区推荐的视频
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

// ChannelRecommendTypeList 频道页-获取频道类型推荐的视频
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
	// 频道ID
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定频道")
		c.ServeJSON()
	}
	// 每页数量
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

// VideoInfo 获得视频详情
func (c *VideoController) VideoInfo()  {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = ReturnError(4005,"必须指定视频ID且id为整数")
		c.ServeJSON()
	}
	video,err := models.GetVideoInfo(id)
	if err == nil {
		c.Data["json"] = SuccessJsonStruct{Items: video}
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"没有相关内容")
		c.ServeJSON()
	}

}
// EpisodesList 获得视频剧集列表
func (c *VideoController) EpisodesList()  {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["json"] = ReturnError(4005,"必须指定视频ID且id为整数")
		c.ServeJSON()
	}
	num, episodes,err := models.GetVideoEpisodesList(id)
	if err == nil {
		c.Data["json"] = SuccessJsonStruct{Items: episodes,Count: num}
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"没有相关内容")
		c.ServeJSON()
	}
}


func (c *VideoController) Save()  {
	// 标题
	title := c.GetString("title")
	if title == "" {
		c.Data["json"] = ReturnError(4011,"视频标题不能为空")
		c.ServeJSON()
	}
	// 子标题
	subTitle := c.GetString("subTitle","")
	// 剧集标题
	episodeTitle := c.GetString("episodeTitle")
	if episodeTitle == "" {
		c.Data["json"] = ReturnError(4011,"剧集标题不能为空")
		c.ServeJSON()
	}
	// 图片
	img := c.GetString("img","")
	// 图片1
	img1 := c.GetString("img1","")
	// 获取频道ID
	channelId, err := c.GetInt("channelId")
	if err != nil {
		c.Data["json"] = ReturnError(4012,"视频ID必须填写")
		c.ServeJSON()
	}
	// 获取频道ID
	regionId, err := c.GetInt("regionId")
	if err != nil {
		c.Data["json"] = ReturnError(4013,"视频地区ID必须填写")
		c.ServeJSON()
	}
	// 获取频道ID
	typeId, err := c.GetInt("typeId")
	if err != nil {
		c.Data["json"] = ReturnError(4014,"视频类型ID必须填写")
		c.ServeJSON()
	}
	// 获取频道ID
	playUrl := c.GetString("playUrl")
	if playUrl == "" {
		c.Data["json"] = ReturnError(4015,"视频播放地址必须填写")
		c.ServeJSON()
	}

	// 获取频道ID
	userId, err := c.GetInt("userId")
	if err != nil {
		c.Data["json"] = ReturnError(4016,"视频用户ID必须填写")
		c.ServeJSON()
	}
	aliyunVideoId := c.GetString("aliyunVideoId","")
	time := time.Now().Unix()

	if videoId,err := models.SaveVideo(title,subTitle,img,img1,channelId,regionId,typeId,userId,time); err == nil {
		if err := models.SaveVideoEpisodes(int(videoId),episodeTitle,playUrl,time,aliyunVideoId);err == nil {
			c.Data["json"] = SuccessJsonStruct{Msg: "添加成功"}
			c.ServeJSON()
		}
	}
	c.Data["json"] = ReturnError(4017,"添加失败")
	c.ServeJSON()
}


func (c *VideoController) Search() {
	//获取搜索关键字
	keyword := c.GetString("keyword")
	//获取翻页信息
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	if keyword == "" {
		c.Data["json"] = ReturnError(4001, "关键字不能为空")
		c.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	sort := []map[string]string{map[string]string{"id": "desc"}}
	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"match": map[string]interface{}{
					"title": keyword,
				},
			},
		},
	}
	res := es.EsSearch("fyouku_video", query, offset, limit, sort)
	var data []models.Video
	for _, v := range res.Hits {
		var itemData models.Video
		err := json.Unmarshal([]byte(v.Source), &itemData)
		if err == nil {
			data = append(data, itemData)
		}
	}
	c.Data["json"] = ReturnSuccess(0, "success", data, int64(len(data)))
	c.ServeJSON()

}


//生成测试视频数据
func (c *VideoController) TestData() {
	//var i = 1
	//for {
	//	i++
	//	rand.Seed(time.Now().UnixNano())
	//	uidi := rand.Intn(10)
	//	uid := uidi + 10
	//	models.SaveVideo(strconv.Itoa(i)+"鸣人柯南一护路飞由诺阿斯塔"+strconv.Itoa(i), "蜡笔小新樱桃小丸子", 1, 2, 2, "/static/video/coverr-sparks-of-bonfire-1573980240958.mp4", uid, "")
	//	i++
	//	c.Data["json"] = SuccessJsonStruct{Msg: " create Test data，title："+"鸣人柯南一护路飞由诺阿斯塔"+strconv.Itoa(i)}
	//	c.ServeJSON()
	//}
}

//导入ES脚本
func (c *VideoController) SendEs() {
	_, data, _ := models.GetAllList()
	for _, v := range data {

		body := map[string]interface{}{
			"id":                   v.Id,
			"title":                v.Title,
			"sub_title":            v.SubTitle,
			"add_time":             v.AddTime,
			"img":                  v.Img,
			"img1":                 v.Img1,
			"episodes_count":       v.EpisodesCount,
			"is_end":               v.IsEnd,
			"channel_id":           v.ChannelId,
			"status":               v.Status,
			"region_id":            v.RegionId,
			"type_id":              v.TypeId,
			"episodes_update_time": v.EpisodesUpdateTime,
			"comment":              v.Comment,
			"user_id":              v.UserId,
			"is_recommend":         v.IsRecommend,
		}
		es.EsAdd("fyouku_video", "video-"+strconv.Itoa(v.Id), body)
	}
	c.Data["json"] = SuccessJsonStruct{Msg: " update_es"}
	c.ServeJSON()
}



