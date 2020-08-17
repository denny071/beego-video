package models

import (
	"encoding/json"
	"fmt"
	"fyoukuApi/services/es"
	redisClient "fyoukuApi/services/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

// Video struct
type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int64
	Comment            int
	UserId             int
	IsRecommend        int
}

// VideoData struct
type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
	Comment       int
}

//Episodes 结构
type Episodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	AliyunVideoId string
}

// 初始化函数
func init() {
	orm.RegisterModel(new(Video))
}

// GetChannelHotList 获得频道热门列表
func GetChannelHotList(channelId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	qs := o.QueryTable("video")
	qs.Filter("status", 1)
	qs.Filter("is_hot", 1)
	qs.Filter("channel_id", channelId)
	qs.OrderBy("-episodes_update_time")
	qs.Limit(9)
	num, err := qs.All(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end","Comment")
	return num, videos, err
}

// GetChannelRecommendRegionList 获得频道推荐地区列表
func GetChannelRecommendRegionList(channelId int, regionId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	qs := o.QueryTable("video")
	qs.Filter("status", 1)
	qs.Filter("region_id", regionId)
	qs.Filter("channel_id", channelId)
	qs.OrderBy("-episodes_update_time")
	qs.Limit(9)
	num, err := qs.All(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	return num, videos, err
}

// GetChannelRecommendTypeList 获得频道推荐类型列表
func GetChannelRecommendTypeList(channelId int, typeId int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	qs := o.QueryTable("video")
	qs.Filter("status", 1)
	qs.Filter("type_id", typeId)
	qs.Filter("channel_id", channelId)
	qs.OrderBy("-episodes_update_time")
	qs.Limit(9)
	num, err := qs.All(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return num, videos, err
}

// GetChannelVideoList 获得频道视频列表
func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video

	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 { // 过滤地区
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 { // 过滤类型
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" { // 是否完结
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 0)
	}
	if sort == "episodesUpdateTime" { // 按照剧集更新时间倒序
		qs = qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" { // 安装评论数评论倒序
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" { // 安装添加时间排序倒序
		qs = qs.OrderBy("-add_time")
	} else { // 默认安装添加时间倒序
		qs = qs.OrderBy("-add_time")
	}
	// 统计数据数量
	num, _ := qs.All(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.All(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return num, videos, err
}

// GetChannelVideoListEs  get video list
func GetChannelVideoListEs(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []Video, error) {
	query := make(map[string]interface{})
	bools := make(map[string]interface{})
	var must []map[string]interface{}
	must = append(must, map[string]interface{}{"term": map[string]interface{}{
		"channel_id": channelId,
	}})
	must = append(must, map[string]interface{}{"term": map[string]interface{}{
		"status": 1,
	}})
	if regionId > 0 {
		must = append(must, map[string]interface{}{"term": map[string]interface{}{
			"region_id": regionId,
		}})
	}
	if typeId > 0 {
		must = append(must, map[string]interface{}{"term": map[string]interface{}{
			"type_id": typeId,
		}})
	}
	if end == "n" {
		must = append(must, map[string]interface{}{"term": map[string]interface{}{
			"is_end": 0,
		}})
	} else if end == "y" {
		must = append(must, map[string]interface{}{"term": map[string]interface{}{
			"is_end": 1,
		}})
	}
	bools["must"] = must
	query["bool"] = bools

	sortData := []map[string]string{map[string]string{"add_time": "desc"}}
	if sort == "episodesUpdateTime" {
		sortData = []map[string]string{map[string]string{"episodes_update_time": "desc"}}
	} else if sort == "comment" {
		sortData = []map[string]string{map[string]string{"comment": "desc"}}
	} else if sort == "addTime" {
		sortData = []map[string]string{map[string]string{"add_time": "desc"}}
	}
	res := es.EsSearch("fyouku_video", query, offset, limit, sortData)
	total := res.Total.Value
	var data []Video
	for _, v := range res.Hits {
		var itemData Video
		err := json.Unmarshal([]byte(v.Source), &itemData)
		if err == nil {
			data = append(data, itemData)
		}
	}
	return int64(total), data, nil
}

// GetUserVideo  get user's video by user id
func GetUserVideo(uid int) (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	qs := o.QueryTable("video")
	qs = qs.Filter("user_id", uid)
	qs = qs.OrderBy("-add_time")
	num, err := qs.All(&videos, "id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end")

	return num, videos, err
}

// GetVideoInfo get video information by video id
func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	qs := o.QueryTable("video")
	qs = qs.Filter("id", videoId)
	err := qs.One(&video)
	return video, err
}

// RedisGetVideoInfo
func RedisGetVideoInfo(videoId int) (Video, error) {
	var video Video
	conn := redisClient.PoolConnect()
	defer conn.Close()
	// define redis key
	redisKey := "video:id:" + strconv.Itoa(videoId)
	// judge whether redis exists or not
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		res, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(res, &video)
	} else {
		o := orm.NewOrm()
		qs := o.QueryTable("video")
		qs = qs.Filter("id", videoId)
		if err = qs.One(&video); err == nil {
			if _, err := conn.Do("hmset", redis.Args{redisKey}.AddFlat(video)...); err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return video, err
}

// GetVideoEpisodesList get video episodes list by video id
func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	qs := o.QueryTable("video_episodes")
	qs = qs.Filter("video_id", videoId)
	qs = qs.OrderBy("-num")
	num, err := qs.All(&episodes, "id", "title", "add_time", "num", "play_url", "comment")
	return num, episodes, err
}

// 增加redis缓存 - 获取视频剧集列表
func RedisGetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	var (
		episodes []Episodes
		num      int64
		err      error
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()

	redisKey := "video:episodes:videoId:" + strconv.Itoa(videoId)
	// judge whether redis key exists or not
	exists, err := redis.Bool(conn.Do("exists", redisKey))
	if exists {
		num, err = redis.Int64(conn.Do("llen", redisKey))
		if err == nil {
			values, _ := redis.Values(conn.Do("lrange", redisKey, "0", "-1"))
			var episodesInfo Episodes
			for _, v := range values {
				if err = json.Unmarshal(v.([]byte), &episodesInfo); err == nil {
					episodes = append(episodes, episodesInfo)
				}
			}
		}
	} else {
		o := orm.NewOrm()
		qs := o.QueryTable("video_episodes")
		qs = qs.Filter("video_id", videoId)
		qs = qs.OrderBy("num")
		num, err = qs.All(&episodes, "id", "title", "add_time", "num", "play_url", "comment", "aliyun_video_id")
		if err == nil {
			for _, v := range episodes {
				jsonValue, err := json.Marshal(v)
				if err == nil {
					conn.Do("rpush", redisKey, jsonValue)
				}
			}
			conn.Do("expire", redisKey, 86400)
		}
	}
	return num, episodes, err
}

// GetChannelTop get channel top by channel_id
func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	qs := o.QueryTable("video")
	qs = qs.Filter("status", 1)
	qs = qs.Filter("channel_id", channelId)
	qs = qs.OrderBy("-comment")
	qs = qs.Limit(10)
	num, err := qs.All(&videos, "id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end")
	return num, videos, err
}

func RedisGetChannelTop(channelId int) (int64, []VideoData, error) {
	var (
		videos []VideoData
		num    int64
	)
	conn := redisClient.PoolConnect()
	defer conn.Close()
	// define redis key
	redisKey := "video:top:channel:channelId:" + strconv.Itoa(channelId)
	// judge whether key exists or not
	exists, err := redis.Bool(conn.Do("exists" + redisKey))
	if exists {
		num = 0
		res, _ := redis.Values(conn.Do("zrevrange", redisKey, "0", "10", "WITHSCORES"))
		for k, v := range res {
			fmt.Println(string(v.([]byte)))
			if k%2 == 0 {
				videoId, err := strconv.Atoi(string(v.([]byte)))
				videoInfo, err := RedisGetVideoInfo(videoId)
				if err == nil {
					var videoDataInfo VideoData
					videoDataInfo.Id = videoInfo.Id
					videoDataInfo.Img = videoInfo.Img
					videoDataInfo.Img1 = videoInfo.Img1
					videoDataInfo.IsEnd = videoInfo.IsEnd
					videoDataInfo.SubTitle = videoInfo.SubTitle
					videoDataInfo.Title = videoInfo.Title
					videoDataInfo.AddTime = videoInfo.AddTime
					videoDataInfo.Comment = videoInfo.Comment
					videoDataInfo.EpisodesCount = videoInfo.EpisodesCount
					videos = append(videos, videoDataInfo)
					num++
				}
			}
		}
	} else {
		o := orm.NewOrm()
		qs := o.QueryTable("video")
		qs = qs.Filter("status", 1)
		qs = qs.Filter("channel_id", channelId)
		qs = qs.OrderBy("-comment")
		qs = qs.Limit(10)
		num, err = qs.All(&videos, "id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end")
		if err == nil {
			// save redis
			for _, v := range videos {
				conn.Do("zadd", redisKey, v.Comment, v.Id)
			}
			conn.Do("expire", redisKey, 86400*30)
		}
	}
	return num, videos, err
}

func SaveVideo(title string, subTitle string, channelId int, regionId int, typeId int, playUrl string, user_id int, aliyunVideoId string) error {
	var video Video
	o := orm.NewOrm()
	time := time.Now().Unix()
	video.Title = title
	video.SubTitle = subTitle
	video.AddTime = time
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.ChannelId = channelId
	video.Status = 1
	video.RegionId = regionId
	video.TypeId = typeId
	video.EpisodesUpdateTime = time
	video.Comment = 0
	video.UserId = user_id
	videoId, err := o.Insert(&video)
	if err == nil {
		if aliyunVideoId != "" {
			playUrl = ""
		}
		_, err = o.Raw("INSERT INTO video_episodes (title,add_time,num,video_id,play_url,statusmcomment,aliyun_video_id) VALUES (?,?,?,?,?,?,?,?)",
			subTitle,time,1,videoId,playUrl,1,0,aliyunVideoId).Exec()
	}
	return err
}
// SaveAliyunVideo save aliyun video
func SaveAliyunVideo(videoId string, log string) error {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO aliyun_video (video_id, log, add_time) VALUES (?,?,?)",videoId,log, time.Now().Unix()).Exec()
	fmt.Println(err)
	return err
}

// GetAllList get all video list
func GetAllList() (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	qs := o.QueryTable("video")
	num, err := qs.All(&videos)
	return num, videos, err
}