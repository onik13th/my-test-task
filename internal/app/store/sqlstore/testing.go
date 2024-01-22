package sqlstore

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"testing"
)

//func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) { // в аргументе функции передается название таблицы
//	t.Helper()
//
//	db, err := sql.Open("postgres", databaseURL)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if err := db.Ping(); err != nil {
//		t.Fatal(err)
//	}
//
//	return db, func(tables ...string) {
//		if len(tables) > 0 {
//			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
//		}
//
//		db.Close()
//	}
//}

func TestDB(t *testing.T, databaseURL string) (*gorm.DB, func(...string)) {
	t.Helper()

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		if err = sqlDB.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
