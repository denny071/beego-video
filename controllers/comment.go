package controllers

import (
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) List() {
	c.Data["json"] = SuccessJsonStruct{Msg: "List"}
	c.ServeJSON()
}

func (c *CommentController) Save() {
	c.Data["json"] = SuccessJsonStruct{Msg: "Save"}
	c.ServeJSON()
}

//生成评论数据
func (c *CommentController) TestData() {
	//var i = 1
	//for {
	//	i++
	//	rand.Seed(time.Now().UnixNano())
	//	uidi := rand.Intn(10)
	//	uid := uidi + 10
	//
	//	models.SaveComment(strconv.Itoa(i)+"博人智商很高，在忍者学校成绩都满分。悟性也都是很高螺旋丸很快就能掌握，但不知道为什么博人在战斗的时候总分不清情况。什么时候该打，什么时候该跑总是表现得一股脑。无论敌人什么实力总是要向前冲，一点也表现不出博人的战术分析。按理说博人比鸣人更有理智，但表现得总是差强人意。我偶尔就看的十分尴尬。"+strconv.Itoa(i), uid, 1, 1)
	//	i++
	//	fmt.Println(i)
	//}
}
