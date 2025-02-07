package processors

import (
	"context"
	"errors"
	"web-server/internals/app/db"
	"web-server/internals/app/models"
)

type UsersProcessor struct {
	storage *db.UsersStorage
}

func NewUsersProcessor(storage *db.UsersStorage) *UsersProcessor {
	processor := new(UsersProcessor)
	processor.storage = storage
	return processor
}

func (processor *UsersProcessor) CreateUser(ctx context.Context, user models.User) error {
	if user.Name == "" {
		return errors.New("name can't be empty")
	}
	return processor.storage.CreateUser(ctx, user)
}

func (processor *UsersProcessor) FindUser(ctx context.Context, id int64) (models.User, error) {
	user := processor.storage.GetUserById(ctx, id)

	if user.Id != id {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (processor *UsersProcessor) ListUsers(ctx context.Context, nameFilter string) ([]models.User, error) {
	return processor.storage.GetUsersList(ctx, nameFilter), nil
}
