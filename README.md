# beego_video 
> 学习beego框架
```go

├── conf
│   └── app.conf        # 配置文件
├── controllers
│   ├── aliyun.go       # 阿里云OSS控制器
│   ├── barrage.go      # 弹幕控制器
│   ├── base.go         # 基础数据控制器
│   ├── comment.go      # 评论控制器
│   ├── common.go       # 公共模块控制器
│   ├── message.go      # 消息控制器
│   ├── top.go          # 排行榜控制器
│   ├── user.go         # 用户控制器
│   └── video.go        # 视频控制
├── go.mod              # mod 包
├── go.sum              # mod 依赖管理
├── main.go             # 入口
├── models             
│   ├── Comment.go      # 评论模型
│   ├── advert.go       # 广告模型
│   ├── barrage.go      # 弹幕模型
│   ├── base.go         # 基础数据模型
│   ├── message.go      # 消息模型
│   ├── user.go         # 用户模型
│   └── video.go        # 视频模型
├── routers
│   └── router.go       # 路由
├── services
│   ├── es              # elticssearch 服务
│   ├── mp              # rabbitmq 服务
│   └── redis           # reidis 服务
└── tests
    └── default_test.go # 默认测试示例

```