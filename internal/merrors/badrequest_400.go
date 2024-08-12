package merrors

import (
	"net/http"

	"motorq-assignment/internal/utils"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            Bad Request ERROR 400                            */
/* -------------------------------------------------------------------------- */
func BadRequest(ctx *gin.Context, err string) {
	var res utils.BaseResponse
	var smerror utils.Error
	errorCode := http.StatusBadRequest

	smerror.Code = errorCode
	smerror.Type = errorType.BadRequest
	smerror.Message = err

	res.Error = &smerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
