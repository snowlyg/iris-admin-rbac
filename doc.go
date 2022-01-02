/**
git tag -d v0.0.1-alpha1 v0.0.1-alpha2 v0.0.1-alpha3 v0.0.1-alpha4 v0.0.1-alpha5 v0.0.1-alpha6 v0.0.1-alpha7 v0.0.1-alpha8 v0.0.1-alpha9 v0.0.1-alpha10 v0.0.1-alpha11 v0.0.1-alpha12 v0.0.1-alpha13 v0.0.1-alpha14 v0.0.1-alpha15 v0.0.1-alpha16 v0.0.1-alpha17 v0.0.1-alpha18 v0.0.1-alpha19
<h1 align="center">IrisAdminRbac</h1>

[IrisAdminRbac](https://www.github.com/snowlyg/iris-admin-rbac) 项目为一个权鉴模块插件,可以为 [IrisAdmin](https://www.github.com/snowlyg/iris-admin) 项目快速集成权鉴管理API.

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
*/

package doc
