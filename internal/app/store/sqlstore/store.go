package sqlstore

import (
	_ "github.com/lib/pq"
	"github.com/onik13th/my-test-task/internal/app/store"
	"gorm.io/gorm"
)

type Store struct {
	db             *gorm.DB
	userRepository *UserRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
