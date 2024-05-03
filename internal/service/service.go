package service

import "context"

type UserService interface {
	Get(ctx context.Context, id int64) (string, error)
}
