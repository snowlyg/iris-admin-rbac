package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	web_response "github.com/snowlyg/iris-admin/server/web/web_gin/response"
	multi "github.com/snowlyg/multi"
	multi_gin "github.com/snowlyg/multi/gin"
)

func Auth() gin.HandlerFunc {
	verifier := multi_gin.NewVerifier()
	verifier.Extractors = []multi_gin.TokenExtractor{multi_gin.FromHeader, multi_gin.FromQuery} // extract token  from Authorization: Bearer $token and query ?token=
	verifier.ErrorHandler = func(ctx *gin.Context, err error) {
		jwtErr, ok := err.(*jwt.ValidationError)
		if ok && jwtErr.Errors == multi.ValidationErrorExpired {
			web_response.Result(http.StatusPaymentRequired, nil, err.Error(), ctx)
			ctx.Abort()
			return
		}
		web_response.UnauthorizedFailWithMessage(err.Error(), ctx)
		ctx.Abort()
	} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}
