package grpc

import (
	"context"

	"github.com/haffjjj/grpc_learn/server/domain/mathe"
	"github.com/haffjjj/grpc_learn/server/domain/mathe/delivery/grpc/proto"
	"google.golang.org/grpc"
)

type matheHandler struct {
	matheUsecase mathe.Usecase
}

//NewMatheHandler ...
func NewMatheHandler(srv *grpc.Server, matheUsecase mathe.Usecase) {
	handler := &matheHandler{matheUsecase}
	proto.RegisterAddServiceServer(srv, handler)
}

func (m *matheHandler) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result, err := m.matheUsecase.Mutiply(a, b)

	if err != nil {
		return nil, err
	}

	return &proto.Response{Result: result}, nil
}
