package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/iris-admin/server/operation"
	multi "github.com/snowlyg/multi/iris"
)

// OperationRecord 操作日志中间件
func OperationRecord() iris.Handler {
	return func(ctx iris.Context) {
		var body []byte
		var err error
		var disbale string
		var rules []string

		ctx.Request().ParseForm()
		disbale = ctx.Request().Form.Get("operation_record_disbale")
		rule := ctx.Request().Form.Get("operation_record_rules")
		rules = strings.Split(rule, ",")

		// 禁用中间件
		if disbale == "1" {
			ctx.Next()
			return
		}

		contentTyp := ctx.Request().Header.Get("Content-Type")
		// 文件上传过滤body,规则设置了 request 过滤body
		ruleType := arr.NewCheckArrayType(len(rules))
		for _, rule := range rules {
			ruleType.Add(rule)
		}
		if !strings.Contains(contentTyp, "multipart/form-data") || !ruleType.Check("request") {
			body, err = ctx.GetBody()
			if err == nil {
				ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.ResponseWriter().Clone(),
			body:           &bytes.Buffer{},
		}
		ctx.ResetResponseWriter(writer)
		now := time.Now()

		ctx.Next()

		latency := time.Since(now)
		errorMessage := ""
		if ctx.GetErr() != nil {
			errorMessage = ctx.GetErr().Error()
		}

		record := &operation.Oplog{
			Ip:           ctx.RemoteAddr(),
			Method:       ctx.Method(),
			Path:         ctx.Path(),
			Agent:        ctx.Request().UserAgent(),
			Body:         string(body),
			UserID:       multi.GetUserId(ctx),
			ErrorMessage: errorMessage,
			Status:       ctx.GetStatusCode(),
			Latency:      latency,
		}
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

// responseBodyWriter 响应主体 writer
type responseBodyWriter struct {
	context.ResponseWriter
	body *bytes.Buffer
}

// Write 写入
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
