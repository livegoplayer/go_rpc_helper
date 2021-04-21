package middleware

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GetGrpcAccessLoggerOptions(logger *logrus.Logger) []grpc.UnaryServerInterceptor {
	logrusEntry := logrus.NewEntry(logger)
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	list := []grpc.UnaryServerInterceptor{
		grpc_logrus.UnaryServerInterceptor(logrusEntry)}
	return list
}
