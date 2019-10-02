### 目录结构

```

---actions
------api
------routers.go
---conf
------app.ini
------excep
---core
------app
------config
------ginapi
------logging
---helper
------utils
---library
---middlerware
---models
------models.go
---runtime
---service
---store
------gredis
------mongo
------mq
---main.go


```

- actions
 
 路由
 
 请求控制

- conf

配置

- core

框架核心

- helper

全局工具箱、项目应用小工具

- library

三方 库/包 的封装

- middleware

中间件

- models

模型

- runtime

- service

业务逻辑

- store


### 需要的包

```

go get -u -v github.com/gin-gonic/gin               #Gin框架
go get -u -v github.com/go-ini/ini                  #ini配置
go get -u -v github.com/jinzhu/gorm                 #GORM
go get -u -v github.com/dgrijalva/jwt-go            #jwt
go get -u -v github.com/unknwon/com                 #综合工具

#数据存储
go get -u -v github.com/gomodule/redigo/redis       #reids
#go get -u -v go get gopkg.in/mgo.v2                #mongoDB
go get -u -v github.com/globalsign/mgo              #mongoDB社区支持版
go get -u -v github.com/Shopify/sarama              #kafka
go get -u -v github.com/streadway/amqp              #rabbitmq

#请求参数验证，这里用第一个
go get -u -v github.com/astaxie/beego/validation
go get -u -v github.com/go-playground/validator
go get -u -v github.com/asaskevich/govalidator

#热编译bee、realize和gowatch，这里用bee
go get github.com/beego/bee                         #bee run命令是监控的项目
go get -u -v github.com/oxequa/realize
go get -u -v github.com/silenceper/gowatch

#swagger API文档
go get -u -v github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles


```
