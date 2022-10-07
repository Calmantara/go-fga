package message

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// custom code:
// 80 -> BAD_REQUEST
// 99 -> INTERNAL_SERVER_ERROR
// 0 -> SUCCESS

// annotation -> OMITEMPTY -> menghapus json property yang
// value null, atau kosong
type Response struct {
	Code      int        `json:"code"`
	Message   string     `json:"message,omitempty"`
	Error     string     `json:"error,omitempty"`      // nullable
	Data      any        `json:"data,omitempty"`       // nullable
	StartTime *time.Time `json:"start_time,omitempty"` // nullable
}

func ErrorResponseSwitcher(ctx *gin.Context, httpCode int) {
	var response Response
	switch httpCode {
	case http.StatusUnauthorized:
		response = Response{
			Code:    98,
			Message: "unauthorized request",
			Error:   "UNAUTHORIZED",
		}
	default:
		response = Response{
			Code:    99,
			Message: "something went wrong",
			Error:   "INTERNAL_SERVER_ERROR",
		}
	}
	ctx.AbortWithStatusJSON(httpCode, response)
}
