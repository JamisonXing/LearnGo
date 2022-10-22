package store

import (
	myStore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*myStore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*myStore.Book
}

func (m MemStore) Create(book *myStore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Update(book *myStore.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Get(s string) (myStore.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) GetAll() ([]myStore.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}
