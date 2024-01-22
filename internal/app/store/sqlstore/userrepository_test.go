package sqlstore_test

import (
	"github.com/onik13th/my-test-task/internal/app/model"
	"github.com/onik13th/my-test-task/internal/app/store"
	"github.com/onik13th/my-test-task/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	createdUser, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestUserRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	id := 1
	_, err := s.User().FindById(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	createdUser, _ := s.User().Create(model.TestUser(t))
	u, err := s.User().FindById(createdUser.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Remove(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	createdUser, _ := s.User().Create(model.TestUser(t))

	err := s.User().Remove(createdUser.ID)
	assert.NoError(t, err)

	u, _ := s.User().FindById(createdUser.ID)
	assert.Nil(t, u)
}

func TestUserRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	createdUser, _ := s.User().Create(model.TestUser(t))

	updatedUser, err := s.User().Update(&model.User{
		Name:       "Иосиф",
		Surname:    "Сталин",
		Patronymic: "Виссарионович",
	}, createdUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.NotEqual(t, createdUser, updatedUser)
}

func TestUserRepository_FindAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
