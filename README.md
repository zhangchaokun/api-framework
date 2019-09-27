### 目录结构

```

---actions
------routers.go
---conf
------app.ini
------excep
---core
---helper
------utils
---library
---middlerware
---models
------models.go
---runtime
---service
---store
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

go get -u -v github.com/gin-gonic/gin
go get -u -v github.com/go-ini/ini
go get -u -v github.com/jinzhu/gorm

下面是几个常用的验证，这里用第一个
go get -u -v github.com/astaxie/beego/validation
go get -u -v github.com/go-playground/validator
go get -u -v github.com/asaskevich/govalidator

go get -u -v github.com/gomodule/redigo/redis
go get -u -v github.com/dgrijalva/jwt-go
go get -u -v github.com/unknwon/com
go get -u -v github.com/oxequa/realize

go get -u -v github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles



```
