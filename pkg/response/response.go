package response

import "github.com/gin-gonic/gin"

const (
	CodeSuccess       = 200
	CodeInvalidParams = 400
	CodeTokenExpired  = 401
	CodeTokenInvalid  = 402
	CodeServerError   = 500
	CodeNotFound     = 404
	CodeForbidden    = 403
)

var msgMap = map[int]string{
	CodeSuccess:       "ok",
	CodeInvalidParams: "请求参数错误",
	CodeTokenExpired:  "token已过期",
	CodeTokenInvalid:  "token验证失败",
	CodeServerError:   "服务器内部错误",
	CodeNotFound:     "资源不存在",
	CodeForbidden:    "禁止访问",
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: CodeSuccess,
		Msg:  msgMap[CodeSuccess],
		Data: data,
	})
}

func Error(c *gin.Context, httpCode, errorCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errorCode,
		Msg:  msgMap[errorCode],
		Data: data,
	})
}

// 添加新的便捷方法
func BadRequest(c *gin.Context, data interface{}) {
	Error(c, 400, CodeInvalidParams, data)
}

func Unauthorized(c *gin.Context, data interface{}) {
	Error(c, 401, CodeTokenExpired, data)
}

func Forbidden(c *gin.Context, data interface{}) {
	Error(c, 403, CodeForbidden, data)
}

func NotFound(c *gin.Context, data interface{}) {
	Error(c, 404, CodeNotFound, data)
}

func ServerError(c *gin.Context, data interface{}) {
	Error(c, 500, CodeServerError, data)
}