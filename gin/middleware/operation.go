package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/iris-admin/server/operation"
	multi "github.com/snowlyg/multi/gin"
)

func OperationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body []byte
		var err error
		var disbale string
		var rules []string

		ctx.Request.ParseForm()
		disbale = ctx.Request.Form.Get("operation_record_disbale")
		rule := ctx.Request.Form.Get("operation_record_rules")
		rules = strings.Split(rule, ",")

		// 禁用中间件
		if disbale == "1" {
			ctx.Next()
			return
		}

		contentTyp := ctx.Request.Header.Get("Content-Type")
		// 文件上传过滤body,规则设置了 request 过滤body
		ruleType := arr.NewCheckArrayType(len(rules))
		for _, rule := range rules {
			ruleType.Add(rule)
		}
		if !strings.Contains(contentTyp, "multipart/form-data") || !ruleType.Check("request") {
			body, err = ioutil.ReadAll(ctx.Request.Body)
			if err == nil {
				// ioutil.ReadAll 读取数据后重新回写数据
				ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		record := &operation.Oplog{
			Ip:        ctx.ClientIP(),
			Method:    ctx.Request.Method,
			Path:      ctx.Request.URL.Path,
			Agent:     ctx.Request.UserAgent(),
			Body:      string(body),
			UserID:    multi.GetUserId(ctx),
			TenancyId: multi.GetTenancyId(ctx),
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = writer
		now := time.Now()

		ctx.Next()

		latency := time.Since(now)
		record.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = ctx.Writer.Status()
		record.Latency = latency

		responseRuleType := arr.NewCheckArrayType(len(rules))
		for _, rule := range rules {
			responseRuleType.Add(rule)
		}
		//规则设置了 response 过滤响应数据
		if !responseRuleType.Check("response") {
			record.Resp = writer.body.String()
		}

		operation.CreateOplog(record)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
