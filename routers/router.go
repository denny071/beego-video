// @APIVersion 1.0.0
// @Title 仿优酷API接口
// @Description 仿优酷API接口文档
// @Contact 917942168@qq.com
// @TermsOfServiceUrl http://localhost:8099

package routers

import (
	"fmt"
	"fyoukuApi/controllers"
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "127.0.0.1" {
				return true
			}
			fmt.Println("只能本地访问")
			return false
		}),
		// 用户模块
		beego.NSNamespace("/user",
			// 用户登录
			beego.NSRouter("/login", &controllers.UserController{}, "post:LoginDo"),
			// 用户模块
			beego.NSRouter("/register", &controllers.UserController{}, "post:SaveRegister"),
			// 用户视频
			beego.NSRouter("/video", &controllers.UserController{}, "get:Video"),
		),
		// 消息模块
		beego.NSNamespace("/message",
			// 发送消息
			beego.NSRouter("/send", &controllers.MessageController{}, "post:Send"),

		),

		// 阿里云模块
		beego.NSNamespace("/aliyun",
			// 上传视频
			beego.NSRouter("/video", &controllers.AliyunController{}, "post:UploadVideo"),
			// 播放授权
			beego.NSRouter("/auth", &controllers.AliyunController{}, "get:PlayAuth"),
			// 回调地址
			beego.NSRouter("/callback", &controllers.AliyunController{}, "get:Callback"),
		),

		// 弹幕
		beego.NSRouter("/barrage", &controllers.BarrageController{}, "get:List;post:Save"),

		// 基础数据
		beego.NSNamespace("/base",
			// 用户登录
			beego.NSRouter("/region", &controllers.BaseController{}, "get:ChannelRegion"),
			// 用户模块
			beego.NSRouter("/type", &controllers.BaseController{}, "get:ChannelType"),
		),

		// 评论
		beego.NSRouter("/comment", &controllers.CommentController{}, "get:List;post:Save"),

		// 排行榜
		beego.NSNamespace("/top",
			// 用户登录
			beego.NSRouter("/channel", &controllers.TopController{}, "get:ChannelTop"),
			// 用户模块
			beego.NSRouter("/type", &controllers.TopController{}, "get:TypeTop"),
		),


		// 视频
		beego.NSNamespace("/video",
			// 视频列表
			beego.NSRouter("/:id", &controllers.VideoController{}, "get:VideoInfo"),
			// 视频剧集列表
			beego.NSRouter("/:id/episodes", &controllers.VideoController{}, "get:EpisodesList"),
			// 视频列表
			beego.NSRouter("/", &controllers.VideoController{}, "get:ChannelVideo;post:Save"),
			// 广告
			beego.NSRouter("/advert", &controllers.VideoController{}, "get:ChannelAdvert"),
			// 正在热播
			beego.NSRouter("/hot", &controllers.VideoController{}, "get:ChannelHotList"),
			// 搜索
			beego.NSRouter("/search", &controllers.VideoController{}, "get:Search"),
			// 搜索
			beego.NSRouter("/test_data", &controllers.VideoController{}, "get:TestData"),


			// 频道
			beego.NSNamespace("/channel",
				// 推荐
				beego.NSNamespace("/recommend",
					// 地区
					beego.NSRouter("/region", &controllers.VideoController{}, "get:ChannelRecommendRegionList"),

					// 类型
					beego.NSRouter("/type", &controllers.VideoController{}, "get:ChannelRecommendTypeList"),

				),
			),
		),




	)
	beego.AddNamespace(ns)

}
