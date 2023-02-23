package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	web_response "github.com/snowlyg/iris-admin/server/web/web_gin/response"
	multi "github.com/snowlyg/multi/gin"
)

func Auth() gin.HandlerFunc {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader, multi.FromQuery} // extract token  from Authorization: Bearer $token and query ?token=
	verifier.ErrorHandler = func(ctx *gin.Context, err error) {
		if jwtErr, ok := err.(jwt.ValidationError); ok && jwtErr.Errors == jwt.ValidationErrorExpired {
			web_response.Result(http.StatusPaymentRequired, nil, err.Error(), ctx)
			ctx.Abort()
		}
		web_response.UnauthorizedFailWithMessage(err.Error(), ctx)
		ctx.Abort()
	} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}
