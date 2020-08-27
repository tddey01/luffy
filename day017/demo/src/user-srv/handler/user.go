package handler

import (
	tlog "github.com/tddey01/luffy/day017/demo/src/share/utils/log"
	"go.uber.org/zap"
)

// 定义绑定方法
type UserHandler struct {
	logger *zap.Logger
}

// 创建结构体对象
func NewUserHandler() *UserHandler {
	return &UserHandler{
		logger: tlog.Instance().Named("UserHandler"),
	}
}

func (h *userServiceHandler) InstertUser(ctx context.Context, in *InsertUserReq, out *InsertUserResp) error {

}
