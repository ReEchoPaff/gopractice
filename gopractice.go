package gopractice

import "sync"

type BookStore struct {
	Height int64
	Width  int64
}

var (
	bookStore     *BookStore
	bookStoreOnce sync.Once
)

func NewBook(h, w int64) *BookStore {
	bookStoreOnce.Do(func() {
		bookStore = &BookStore{h, w}
	})
	return bookStore
}

func (b *BookStore) ShowBookStoreHeight() int64 {
	return b.Height
}
