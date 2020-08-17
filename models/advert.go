package models

import "github.com/astaxie/beego/orm"

// Advert struct
type Advert struct {
	Id       int
	Title    string
	SubTitle string
	AddTime  int64
	Img      string
	Url      string
}
// initialization model
func init() {
	orm.RegisterModel(new(Advert))
}

// GetChannelAdvert 获得频道广告
func GetChannelAdvert(channelId int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var adverts []orm.Params
	qs := o.QueryTable("advert")
	qs.Filter("state", 1)
	qs.Filter("channel_id", channelId)
	qs.OrderBy("-sort")
	qs.Limit(1)
	num, err := qs.Values(&adverts, "id", "title", "sub_title", "img", "add_time", "url")
	return num, adverts, err

}
