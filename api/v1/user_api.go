package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hurricane5250/go-project/service"
	"strconv"
)

type userApi struct {
	*Api
	*service.UserService
}

var UserApi = userApi{
	Api:         BaseApi,
	UserService: service.NewUserService(),
}

// GetById 文章详情
func (r userApi) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		r.Failed(ctx, ParamError, "id不能为空")
		return
	}
	user, err := r.UserService.GetById(ctx, id)
	if err != nil {
		r.Failed(ctx, Failed, err.Error())
	} else {
		r.Success(ctx, "ok", user)
	}
	return
}
