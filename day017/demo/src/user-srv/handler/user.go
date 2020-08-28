package handler

import (
	"context"
	pb "github.com/tddey01/luffy/day017/demo/src/share/pb"
	tlog "github.com/tddey01/luffy/day017/demo/src/share/utils/log"
	"github.com/tddey01/luffy/day017/demo/src/user-srv/db"
	"github.com/tddey01/luffy/day017/demo/src/user-srv/emtity"
	"go.uber.org/zap"
	"log"
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

func (c *UserHandler) InstertUser(ctx context.Context, req *pb.InsertUserReq, resp *pb.InsertUserResp) error {
	log.Println("InsertUser")
	// 封装结构体
	user := &emtity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
	// 调用数据方法插入
	insertId, err := db.InserUser(user)
	if err != nil {
		log.Fatal("添加用户错误")
		return err
	}
	resp.Id = int32(insertId)
	return nil
}

// 删除
func (c *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserReq, resp *pb.DeleteUserResp) error {
	log.Println("DeleteUser...")
	err := db.DeleteUser(req.GetId())
	if err != nil {
		log.Fatal("删除用户错误")
		return err
	}
	return nil
}

// 修改
func (c *UserHandler) ModifyUser(ctx context.Context, req *pb.ModifyUserReq, resp *pb.ModifyUserResp) error {
	log.Println("update User", req.GetId())
	//	封装结构体
	user := &emtity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Id:      req.Id,
	}
	err := db.ModifyUser(user)
	if err != nil {
		log.Fatal("修改用户错误")
		return err
	}
	return nil
}

// 查询
func (c *UserHandler) SelecttUser(ctx context.Context, req *pb.SelectUserReq, resp *pb.SelectUserResp) error {
	log.Println("Select User")
	user, err := db.SelectUserById(req.Id)
	if err != nil {
		log.Fatal("查询用户错误")
		return err
	}
	if user != nil {
		resp.Users = user.ToProtoUser()
	}
	return nil
}
