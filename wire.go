//go:build wireinject

package main

import (
	"github.com/asynccnu/be-feedback_help/grpc"
	"github.com/asynccnu/be-feedback_help/ioc"
	"github.com/asynccnu/be-feedback_help/pkg/grpcx"
	"github.com/asynccnu/be-feedback_help/repository"
	"github.com/asynccnu/be-feedback_help/repository/cache"
	"github.com/asynccnu/be-feedback_help/repository/dao"
	"github.com/asynccnu/be-feedback_help/service"
	"github.com/google/wire"
)

func InitGRPCServer() grpcx.Server {
	wire.Build(
		ioc.InitGRPCxKratosServer,
		grpc.NewFeedbackHelpServiceServer,
		service.NewFeedbackHelpService,
		repository.NewFeedbackHelpHelpRepository,
		dao.NewFeedbackHelpGormDao,
		cache.NewFeedbackHelpRedisCache,
		// 第三方
		ioc.InitEtcdClient,
		ioc.InitRedis,
		ioc.InitDB,
		ioc.InitLogger,
	)
	return grpcx.Server(nil)
}
