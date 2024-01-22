package store

type Store interface {
	User() UserRepository
}

// интерфейс "Store", который определяет функционал для работы с хранилищем данных.
// этот интерфейс содержит только один метод User(), который возвращает экземпляр UserRepository.
