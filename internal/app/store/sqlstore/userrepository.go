package sqlstore

import (
	"errors"
	"github.com/onik13th/my-test-task/internal/app/model"
	"github.com/onik13th/my-test-task/internal/app/store"
	"gorm.io/gorm"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := r.store.db.Create(&u).Error; err != nil {

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.First(u, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, store.ErrRecordNotFound
		}
		// Другие ошибки
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Remove(id int) error {
	u := &model.User{}
	result := r.store.db.Where("id = ?", id).Delete(u)
	if result.Error != nil {
		return result.Error // Ошибка связанная с бд
	}
	if result.RowsAffected == 0 {
		return store.ErrRecordNotFound // Нет записи для удаления
	}

	return nil
}

func (r *UserRepository) Update(u *model.User, id int) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	existingUser := &model.User{}
	if err := r.store.db.Where("id = ?", id).First(existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	if err := r.store.db.Model(existingUser).Omit("id").Updates(u).Error; err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (r *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.store.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
//	u := &model.User{}
//	if err := r.store.db.Where("email = ?", email).First(u).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, fmt.Errorf("user with email '%s' not found", email)
//		}
//		return nil, err
//	}
//	return u, nil
//}
