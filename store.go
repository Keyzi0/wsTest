package main

// эмуляция хранилища
type Store struct {
	data map[int64]PoloniexMsg
}

// возвращаем инициализированный объект типа Store
func NewStore() Store {
	data := make(map[int64]PoloniexMsg)
	return Store{
		data: data,
	}
}

// методы Get, Set, GetAll как пример реализации CRUD(без D :)
func (s *Store) Set(rec PoloniexMsg) {
	s.data[rec.pairID] = rec
}

func (s *Store) Get(id int64) PoloniexMsg {
	return s.data[id]
}

func (s *Store) GetAll() map[int64]PoloniexMsg {
	return s.data
}