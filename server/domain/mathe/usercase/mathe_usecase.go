package usecase

import (
	"github.com/haffjjj/grpc_learn/server/domain/mathe"
)

type matheUsecase struct{}

//NewMatheUsecase ...
func NewMatheUsecase() mathe.Usecase {
	return &matheUsecase{}
}

func (m *matheUsecase) Mutiply(a, b int64) (int64, error) {
	result := a * b

	return result, nil
}
