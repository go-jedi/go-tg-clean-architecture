package repository

import "context"

type UserRepository interface {
	Get(ctx context.Context, id int64) (string, error)
}
