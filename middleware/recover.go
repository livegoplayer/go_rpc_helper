package middleware

import (
	"errors"
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	myLogger "github.com/livegoplayer/go_logger/logger"
	"google.golang.org/grpc"
	"runtime/debug"
)

func GetGrpcErrHandler(isDebug bool) grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(GetErrorLog(isDebug)))
}

func ErrorLog(p interface{}) (err error) {
	err = errors.New(fmt.Sprintf("%v", p))
	myLogger.Error(err.Error())
	return err
}

func GetErrorLog(isDebug bool) func(p interface{}) (err error) {
	return func(p interface{}) (err error) {
		if isDebug {
			debug.PrintStack()
		}
		return ErrorLog(p)
	}
}
