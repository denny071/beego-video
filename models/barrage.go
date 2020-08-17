package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Barrage struct {
	Id          int
	Content     string
	CurrentTime int
	AddTime     int64
	UserId      int
	Status      int
	EpisodesId  int
	VideoId     int
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func init() {
	orm.RegisterModel(new(Barrage))
}

func BarrageList(episodesId int, startTime int, endTime int) (int64, []BarrageData, error) {
	o := orm.NewOrm()
	var barrages []BarrageData
	qs := o.QueryTable("barrage")
	qs = qs.Filter("status", 1)
	qs = qs.Filter("episodes_id", episodesId)
	qs = qs.Filter("current_time__gte", startTime)
	qs = qs.Filter("current_time__lt", endTime)
	qs = qs.OrderBy("current_time")
	num, err := qs.All(&barrages, "id", "content", "current_time")
	return num, barrages, err
}
// SaveBarrage save barrage
func SaveBarrage(episodesId int, videoId int, currentTime int, userId int, content string) error {
	_, err := orm.NewOrm().Insert(&Barrage{
		Content:     content,
		CurrentTime: currentTime,
		AddTime:     time.Now().Unix(),
		UserId:      userId,
		Status:      1,
		EpisodesId:  episodesId,
		VideoId:     videoId,
	})
	return err
}
