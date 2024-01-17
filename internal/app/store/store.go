package store

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	config *Config
	db     *gorm.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := gorm.Open(postgres.Open(s.config.DatabaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
