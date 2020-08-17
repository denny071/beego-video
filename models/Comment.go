package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id          int
	Comment     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	Video       int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	qs := o.QueryTable("comment")
	qs = qs.Filter("status", 1)
	qs = qs.Filter("episodes_id", episodesId)
	num, _ := qs.Count()
	qs = qs.Offset(offset)
	qs = qs.Limit(limit)
	qs = qs.OrderBy("-add_time")
	_, err := qs.All(&comments, "id", "content", "add_time", "user_id", "stamp", "praise_count", "episodes_id")

	return num, comments, err
}

//
//func SaveComment(content string, uid int, episodesId int, videoId int) error {
//	o := orm.NewOrm()
//	comment := Comment{
//		Comment: content,
//		UserId:  uid,
//		EpisodesId: episodesId,
//		Video: videoId,
//		Stamp: 0,
//		Status: 1,
//		AddTime: time.Now().Unix(),
//	}
//	_, err := o.Insert(comment)
//	if err == nil {
//		// 修改视频总评论数
//		o.QueryTable("video").Filter("id",videoId).Update(orm.Params{"comment":orm.ColValue(orm.ColAdd,1)})
//		// 修改视频剧集的评论数
//		o.QueryTable("video_episodes").Filter("id",episodesId).Update(orm.Params{"comment":orm.ColValue(orm.ColAdd,1)})
//
//		// 更新redis排行榜 - 通过MQ来实现
//		videoObj := map[string]int {
//			"VideoId":videoId,
//		}
//		videoJson, _ := json.Marshal(videoObj)
//
//	}
//
//}