package controllers

import (
	"beego-video/models"
	"github.com/astaxie/beego"
)


type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int             `json:"episodesId"`
}

type CommentController struct {
	beego.Controller
}

func (c *CommentController) List() {
	// 获得剧集ID
	episodesId, _ := c.GetInt("episodesId")
	// 获得页码信息
	limit,_ := c.GetInt("limit")
	offset,_ := c.GetInt("offset")

	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001,"必定指定视频剧集")
		c.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId,offset,limit)
	if err == nil {
		var data []CommentInfo
		var commentInfo CommentInfo

		// 获取uid channel
		uidChan := make(chan int, 12)
		closeChan := make(chan bool,5)
		resChan := make(chan models.UserInfo,12)
		// 把获取到的uid放到channel中
		go func() {
			for _, v := range comments {
				uidChan <- v.UserId
			}
			close(uidChan)
		}()
		// 处理uidChannel中的信息
		for i := 0; i < 5; i++ {
			go chanGetUserInfo(uidChan, resChan, closeChan)
		}
		// 判断是否执行完成，信息聚合
		go func() {
			for i := 0; i < 5; i++ {
				<-closeChan
			}
			close(resChan)
			close(closeChan)
		}()

		userInfoMap := make(map[int]models.UserInfo)
		for r := range resChan {
			userInfoMap[r.Id] = r
		}

		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DataFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.EpisodesId = v.EpisodesId
			// 获取用户信息
			commentInfo.UserInfo, _ = userInfoMap[v.UserId]
			data = append(data, commentInfo)
		}
		c.Data["json"] = ReturnSuccess(0,"success",data,num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004,"没有相关内容")
		c.ServeJSON()
	}
}

func chanGetUserInfo(uidChan chan int, resChan chan models.UserInfo, closeChan chan bool)  {
	for uid := range uidChan {
		res, err := models.GetUserInfo(uid)
		if err == nil {
			resChan <- res
		}
	}
	closeChan <- true
}

func (c *CommentController) Save() {
	content := c.GetString("content")
	uid,_ := c.GetInt("uid")
	episodesId, _ := c.GetInt("episodesId")
	videoId, _ := c.GetInt("videoId")
	if content == "" {
		c.Data["json"] = ReturnError(4001,"内容不能为空")
		c.ServeJSON()
	}
	if uid == 0 {
		c.Data["json"] = ReturnError(4001,"请先登录")
		c.ServeJSON()
	}
	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定评论剧集ID")
		c.ServeJSON()
	}
	if videoId == 0 {
		c.Data["json"] = ReturnError(4001,"必须指定视频ID")
		c.ServeJSON()
	}
	err := models.SaveComment(content,uid,episodesId,videoId)
	if err == nil {
		c.Data["json"] = ReturnSuccess(0,"success","",1)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(5000,err)
		c.ServeJSON()
	}

}
