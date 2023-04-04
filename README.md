
# go-zero单体服务 - 脚手架

## goctl 命令行使用

```
- 创建项目 - goctl api new 【项目名称】 -dir . -style go_zero
- 生成模型 - goctl model mysql ddl -src sql/【数据表】.sql -dir ./model/【文件夹】
- 生成接口 - goctl api go -api app/apictl/app.api -dir ./app -style go_zero
```

## 目录结构
```
├── README.md
├── app 【程序】
│ ├── app.api
│ ├── app.go
│ ├── etc
│ └── internal
├── common 【公共】
├── go.mod
├── go.sum
├── model 【模型】
│ └── demo_user
└── sql 【sql】
└── demo_user.sql
```

```
日志 中间件 数据库
```