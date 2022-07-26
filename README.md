# <center>Ashi</center>

### <center>基于gin的一个web框架，封装了gorm，redis，jwt验证，logrus</center>

### 结构

```
ashi
│  users.sql 基础用户表
│  config.ini 配置文件
│  Dockerfile 
│  docker-compose.yml 
│  main.go
├─controller 控制器
│      BaseController.go 基础控制器
│      UserController.go 用户模块控制器 
├─middleware 中间件
│      auth.go token验证中间件
├─model 模型
│      base.go 基础模型
│      user.go 用户模型
│      character.go 角色模型
│      character_pic.go 角色立绘模型
├─router 路由
│      api.go 
├─runtime
│  │─logs
│  └─upload 上传文件路径
├─service 服务
│      user.go 用户模块服务层
├─setting
│      config.go 配置信息
│      setup.go 项目启动
└─utils
        help.go 封装的小工具函数
        jwt.go jwt基础
        log.go 日志工具
        md5.go md5工具
        random.go 随机字符串工具
        upload.go 上传文件工具
```