### 目录结构

```

---actions
------routers.go
---conf
------app.ini
------excep
---core
---middlerware
---models
------models.go
---runtime
---service
---store
---utils
---main.go


```

- actions
 
 路由
 
 请求控制

- conf

配置

- core

框架核心

- middleware

中间件

- models

模型

- runtime

- service

业务逻辑

- store

- utils

工具箱


### 需要的包

```

go get -u -v github.com/gin-gonic/gin
go get -u -v github.com/go-ini/ini
go get -u -v github.com/jinzhu/gorm

下面两个是常用的，这里用第一个
go get -u -v github.com/go-playground/validator
go get -u -v github.com/asaskevich/govalidator

go get -u -v github.com/gomodule/redigo/redis
go get -u -v github.com/dgrijalva/jwt-go

go get -u -v github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles



```