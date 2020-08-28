package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/tddey01/luffy/day017/demo/src/share/config"
	pb "github.com/tddey01/luffy/day017/demo/src/share/pb"
	"github.com/tddey01/luffy/day017/demo/src/share/utils/log"
	"github.com/tddey01/luffy/day017/demo/src/user-srv/db"
	"github.com/tddey01/luffy/day017/demo/src/user-srv/handler"
)

func main() {
	logger := log.Init("user")
	// 创建service
	service := micro.NewService(
		micro.Name(config.Namespace+"user"),
		micro.Version("latest"),
	)
	// 初始化 service
	service.Init(
		micro.Action(func(c *cli.Context) {
			logger.Info("user-srv服务运行时间打印...")
			// 初始化db
			db.Init(config.MysqlDNS)
			//	 注册服务
			err := pb.RegisterUserServiceHandler(service.Server(), handler.NewUserHandler(), server.InternalHandler(true))
			if err != nil {
				fmt.Println(err)
			}
		}),
		// 定义服务停止后做的事情
		micro.AfterStop(func() error {
			logger.Info("user-srv服务停止后的打印...")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("user-srv服务启动前的打印...")
			return nil
		}),
	)
	logger.Info("启动user-srv服务....")
	// 启动service
	if err := service.Run(); err != nil {
		logger.Panic(" user-srv服务启动失败")
	}

}
