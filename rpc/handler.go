package main

import (
	"context"
	"github.com/LynnWonder/gin_prac/rpc/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// 返回传进来的内容
	return &api.Response{Message: req.Message}, nil
}
