package api

import (
	"github.com/snowlyg/iris-admin/server/database/orm"
)

type DeleteApiReq struct {
	Path   string `json:"path" form:"path" binding:"required"`
	Method string `json:"method" form:"method" binding:"required"`
}
type AuthorityType struct {
	AuthorityType int `json:"authorityType" form:"authorityType" uri:"authorityType" param:"authorityType"`
	IsMenu        int `json:"isMenu" form:"isMenu" uri:"isMenu" param:"isMenu"`
}

type ReqPaginate struct {
	orm.Paginate
	AuthorityType
	Path        string `json:"path" form:"path"`
	Description string `json:"description" form:"description"`
	ApiGroup    string `json:"apiGroup" form:"apiGroup"`
	Method      string `json:"method" form:"method"`
}
