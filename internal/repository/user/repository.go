package user

import (
	"github.com/go-jedi/go-telegram-clean-architecture/internal/repository"
	"github.com/go-jedi/platform_common/pkg/db"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

/*
const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)
*/
