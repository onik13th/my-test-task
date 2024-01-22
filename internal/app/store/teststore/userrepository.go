package teststore

import (
	"github.com/onik13th/my-test-task/internal/app/model"
	"github.com/onik13th/my-test-task/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	u.ID = len(r.users)
	r.users[u.ID] = u

	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) Remove(id int) error {
	_, ok := r.users[id]
	if !ok {
		return store.ErrRecordNotFound
	}

	delete(r.users, id)

	return nil
}

func (r *UserRepository) Update(u *model.User, id int) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	_, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	r.users[id] = u

	return u, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	var res []*model.User
	m := r.users
	for _, u := range m {
		res = append(res, u)
	}

	return res, nil
}
