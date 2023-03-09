<h1 align="center">IrisAdminRbac</h1>

[IrisAdminRbac](https://www.github.com/snowlyg/iris-admin-rbac) 项目为一个权鉴模块插件,可以为 [IrisAdmin](https://www.github.com/snowlyg/iris-admin) 项目快速集成权鉴管理API.


[![Build Status](https://app.travis-ci.com/snowlyg/iris-admin-rbac.svg?branch=main)](https://app.travis-ci.com/snowlyg/iris-admin-rbac)
[![LICENSE](https://img.shields.io/github/license/snowlyg/iris-admin-rbac)](https://github.com/snowlyg/iris-admin-rbac/blob/main/LICENSE)
[![go doc](https://godoc.org/github.com/snowlyg/iris-admin-rbac?status.svg)](https://godoc.org/github.com/snowlyg/iris-admin-rbac)
[![go report](https://goreportcard.com/badge/github.com/snowlyg/iris-admin-rbac)](https://goreportcard.com/badge/github.com/snowlyg/iris-admin-rbac)
[![Build Status](https://codecov.io/gh/snowlyg/iris-admin-rbac/branch/main/graph/badge.svg)](https://codecov.io/gh/snowlyg/iris-admin-rbac)

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

#### Code snippets

**Code snippets for iris-admin-rbac**
```json
{
  // Place your snippets for go here. Each snippet is defined under a snippet name and has a prefix, body and 
  // description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
  // $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the 
  // same ids are connected.
  // Example:
  // "Print to console": {
  // 	"prefix": "log",
  // 	"body": [
  // 		"console.log('$1');",
  // 		"$2"
  // 	],
  // 	"description": "Log output to console"
  // }
  "Print iris-admin controller": {
    "prefix": "iac",
    "body": [
      "// $1 $2",
      "func $1(ctx *gin.Context) {",
      "  req := &ReqPaginate{}",
      "  if errs := ctx.ShouldBind(&req); errs != nil {",
      "    response.FailWithMessage(errs.Error(), ctx)",
      "    return",
      "  }",
      "  items := &PageResponse{}",
      "  var scopes []func(db *gorm.DB) *gorm.DB",
      "  total, err := items.Paginate(database.Instance(), req.PaginateScope(), scopes...)",
      "  if err != nil {",
      "    response.FailWithMessage(err.Error(), ctx)",
      "    return",
      "  }",
      "  response.OkWithData(response.PageResult{",
      "    List:     items.Item,",
      "    Total:    total,",
      "    Page:     req.Page,",
      "    PageSize: req.PageSize,",
      "  }, ctx)",
      "}",
    ],
    "description": "Print iris-admin controller"
  },
  "Print iris-admin scope": {
    "prefix": "ias",
    "body": [
      "func $1Scope($2 $3) func(db *gorm.DB) *gorm.DB {",
      "  return func(db *gorm.DB) *gorm.DB {",
      "    return db.Where(\"$4 =?\", $2)",
      "  }",
      "}",
    ],
    "description": "Print iris-admin scope"
  },
  "Print iris-admin route group": {
    "prefix": "iag",
    "body": [
      "import (",
      "  \"github.com/gin-gonic/gin\"",
      "  \"github.com/snowlyg/iris-admin-rbac/gin/middleware\"",
      ")",
      "//router.GET(\"/list\", All)",
      "//router.POST(\"/remark/:id\", Remark)",
      "func Group(app *gin.RouterGroup) {",
      "  router := app.Group(\"$1\", middleware.Auth(), middleware.CasbinHandler(), middleware.OperationRecord())",
      "  {",
      "    $2",
      "  }",
      "}",
    ],
    "description": "Print iris-admin route group"
  },
  "Print iris-admin migrate": {
    "prefix": "iami",
    "body": [
      "func GetMigration() *gormigrate.Migration {",
      "  return &gormigrate.Migration{",
      "    // 20211215120700_create_xxxxs_table",
      "    ID: \"$1\",",
      "    Migrate: func(tx *gorm.DB) error {",
      "      return tx.AutoMigrate(&$2{})",
      "    },",
      "    Rollback: func(tx *gorm.DB) error {",
      "      return tx.Migrator().DropTable(\"$3\")",
      "    },",
      "  }",
      "}"
    ],
    "description": "Print iris-admin route group"
  },
  "Print iris-admin model": {
    "prefix": "iam",
    "body": [
      "import (",
      "  \"time\"",
      "  \"gorm.io/gorm\"",
      ")",
      "type $1 struct {",
      "  gorm.Model",
      "  Base$1",
      "}",
      "type Base$1 struct {",
      "  $3",
      "}",
    ],
    "description": "Print iris-admin model"
  },
  "Print iris-admin request": {
    "prefix": "iareq",
    "body": [
      "import (",
      "  \"github.com/gin-gonic/gin\"",
      "  \"github.com/snowlyg/iris-admin/server/database/orm\"",
      "  \"github.com/snowlyg/iris-admin/server/zap_server\"",
      ")",
      "type Request struct {",
      "  Base$1",
      "}",
      "func (req *Request) Request(ctx *gin.Context) error {",
      "  if err := ctx.ShouldBindJSON(req); err != nil {",
      "    zap_server.ZAPLOG.Error(err.Error())",
      "    return orm.ErrParamValidate",
      "  }",
      "  return nil",
      "}",
      "type ReqPaginate struct {",
      "  orm.Paginate",
      "  $2",
      "}"
    ],
    "description": "Print iris-admin request"
  },
  "Print iris-admin response": {
    "prefix": "iares",
    "body": [
      "import (",
      "  \"github.com/gin-gonic/gin\"",
      "  \"github.com/snowlyg/iris-admin/server/database/orm\"",
      "  \"github.com/snowlyg/iris-admin/server/zap_server\"",
      ")",
      "type Response struct {",
      "  orm.Model",
      "  Base$1",
      "}",
      "func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {",
      "  db = db.Model(&$2{})",
      "  if len(scopes) > 0 {",
      "    db.Scopes(scopes...)",
      "  }",
      "  err := db.First(res).Error",
      "  if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {",
      "    zap_server.ZAPLOG.Error(err.Error())",
      "    return err",
      "  }",
      "  return nil",
      "}",
      "// Paginate",
      "type PageResponse struct {",
      "  Item []*Response",
      "}",
      "func (res *PageResponse) Paginate(db *gorm.DB, pageScope func(db *gorm.DB) *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) (int64, error) {",
      "  db = db.Model(&$2{})",
      "  if len(scopes) > 0 {",
      "    db.Scopes(scopes...)",
      "  }",
      "  var count int64",
      "  err := db.Count(&count).Error",
      "  if err != nil {",
      "    zap_server.ZAPLOG.Error(err.Error())",
      "    return 0, err",
      "   }",
      "  db.Scopes(pageScope)",
      "  err = db.Find(&res.Item).Error",
      "  if err != nil {",
      "    zap_server.ZAPLOG.Error(err.Error())",
      "    return 0, err",
      "  }",
      "  return count, nil",
      "}",
      "func (res *PageResponse) Find(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {",
      "  db = db.Model(&$2{})",
      "  if len(scopes) > 0 {",
      "    db.Scopes(scopes...)",
      "  }",
      "  err := db.Find(&res.Item).Error",
      "  if err != nil {",
      "    zap_server.ZAPLOG.Error(err.Error())",
      "    return err",
      "  }",
      "  return nil",
      "}",
    ],
    "description": "Print iris-admin response"
  },
}
```