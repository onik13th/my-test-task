package store

import "github.com/onik13th/my-test-task/internal/app/model"

type UserRepository interface {
	Create(*model.User) (*model.User, error)
	FindAll() ([]*model.User, error)
	FindById(int) (*model.User, error)
	Update(*model.User, int) (*model.User, error)
	Remove(int) error
}
