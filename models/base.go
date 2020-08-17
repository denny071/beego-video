package models

import "github.com/astaxie/beego/orm"

type Region struct {
	Id   int
	Name string
}

type Type struct {
	Id   int
	Name string
}

func GetChannelRegion(channelId int) (int64, []Region, error) {
	o := orm.NewOrm()
	var regions []Region
	qs := o.QueryTable("channel_region")
	qs = qs.Filter("status", 1)
	qs = qs.Filter("channel_id", channelId)
	qs = qs.OrderBy("-sort")
	num, err := qs.All(&regions, "id", "name")
	return num, regions, err
}

func GetChannelType(channelId int) (int64, []Type, error) {
	o := orm.NewOrm()
	var types []Type
	qs := o.QueryTable("channel_type")
	qs = qs.Filter("status", 1)
	qs = qs.Filter("channel_id", channelId)
	qs = qs.OrderBy("-sort")
	num, err := qs.All(&types, "id", "name")
	return num, types, err
}
