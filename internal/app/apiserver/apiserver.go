package apiserver

import (
	"database/sql"
	"github.com/onik13th/my-test-task/internal/app/store/sqlstore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
		}
	}(sqlDB)

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

//Start принимает на вход указатель на объект Config и запускает сервер HTTP на указанном адресе

func newDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err // ошибка подключения к бд
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// newDB открывает новое подключение к базе данных PostgreSQL с
//использованием переданной строки подключения (databaseURL).
//Если подключение успешно установлено, функция возвращает объект базы данных gorm.DB,
//который затем используется для выполнения запросов к базе данных.
