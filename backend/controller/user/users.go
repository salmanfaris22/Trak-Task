package user

import (
	users "todo/domain/user"
	"todo/services"
	"todo/utils/errors"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user users.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		err := errors.NewBadReqstError("Invalid JSON body")
		ctx.JSON(err.Status, err)
		return
	}

	services.CreatUser(user)
}
