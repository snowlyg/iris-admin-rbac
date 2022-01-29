<h1 align="center">IrisAdminRbac</h1>

[IrisAdminRbac](https://www.github.com/snowlyg/iris-admin-rbac) 项目为一个权鉴模块插件,可以为 [IrisAdmin](https://www.github.com/snowlyg/iris-admin) 项目快速集成权鉴管理API.

##### 下载

```sh
  go get -u github.com/snowlyg/iris-admin-rbac@latest
```

##### 简单使用

- gin

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

- iris
  
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

##### 接口说明

- 公共接口 
  
```txt
GET /public/captcha  //验证码
POST /public/admin/login  // 登陆
GET /public/logout  // 退出
GET /public/clean  // 清空所有授权,如果是多点登陆方式
```

- 操作记录
  
```txt
GET /oplog/getOplogList //列表
```

- 管理员

```txt
GET /admin/getAll //列表
GET /admin/getAdmin/:id //列表
POST /admin/createAdmin //添加
PUT /admin/updateAdmin/:id //更新
DELETE /admin/deleteAdmin/:id //删除
GET /admin/profile //登陆用户信息
POST /admin/changeAvatar //更新头像
```

- 接口权限

```txt
GET /api/getList            // 获取Api列表
GET /api/getAll             // 获取所有api
GET /api/getApiById/:id     // 获取单条Api消息
POST /api/createApi         // 创建Api
DELETE /api/deleteApi/:id   // 删除Api
PUT /api/updateApi/:id      // 更新api
DELETE /api/deleteApisByIds // 删除选中api
```

- 授权角色

```txt
GET /authority/getAuthorityList        // 获取角色列表
GET /authority/getAdminAuthorityList   // 获取员工角色列表
GET /authority/getTenancyAuthorityList // 获取商户角色列表
GET /authority/getGeneralAuthorityList // 获取普通用户角色列表
POST /authority/createAuthority        // 创建角色
PUT /authority/updateAuthority/:id     // 更新角色
POST /authority/copyAuthority/:id      // 复制
DELETE /authority/deleteAuthority/:id  // 删除角色
```
