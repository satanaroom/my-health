package user

import "context"

type Service interface {
	Create(ctx context.Context) error
}

type service struct {
}

func NewService() *service {
	return &service{}
}
