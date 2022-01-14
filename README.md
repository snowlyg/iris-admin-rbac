<h1 align="center">IrisAdminRbac</h1>

[IrisAdminRbac](https://www.github.com/snowlyg/iris-admin-rbac) 项目为一个权鉴模块插件,可以为 [IrisAdmin](https://www.github.com/snowlyg/iris-admin) 项目快速集成权鉴管理API.

##### 下载

```sh
  go get -u github.com/snowlyg/iris-admin-rbac@latest
```

#####
- 简单使用(gin)
```go
package main

import (
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
  rbac "github.com/snowlyg/iris-admin-rbac/gin"
)

func main() {
  wi := web_gin.Init()
  v1 := wi.Group("/api/v1")
	{
		rbac.Party(v1)    // 权鉴模块
	}
	web.Start(wi)
}

```

- 简单使用(iris)
```go
package main

import (
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
  rbac "github.com/snowlyg/iris-admin-rbac/iris"
)

func main() {
  wi := web_iris.Init()
  v1 := wi.Group("/api/v1")
	{
		rbac.Party(v1)    // 权鉴模块
	}
	web.Start(wi)
}

```