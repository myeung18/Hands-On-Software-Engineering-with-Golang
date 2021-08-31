package isp

import (
	"context"
	"errors"
)

type Item struct{}

type Fatface interface {
	BatchGetItem(IDs ...int) ([]Item, error)
	BatchGetItemWithContext(ctx context.Context, IDs ...int) ([]Item, error)

	BatchPutItem(items ...Item) error
	BatchPutItemWithContext(ctx context.Context, items ...Item) error

	DeleteItem(ID int) error
	DeleteItemWithContext(ctx context.Context, item Item) error

	GetItem(ID int) (Item, error)
	GetItemWithContext(ctx context.Context, ID int) (Item, error)

	PutItem(item Item) error
	PutItemWithContext(ctx context.Context, item Item) error

	Query(query string, args ...interface{}) ([]Item, error)
	QueryWithContext(ctx context.Context, query string, args ...interface{}) ([]Item, error)

	UpdateItem(item Item) error
	UpdateItemWithContext(ctx context.Context, item Item) error
}

type Cache struct {
	db Fatface
}

func (c Cache) Get(key string) interface{} {
	c.db.GetItem(49)

	return 0
}

func callIt(ft Fatface) {

}

type Value interface {
	Value(key interface{}) interface{}
}

type Monitor interface {
	Done() <-chan struct{}
}
//https://learning.oreilly.com/library/view/hands-on-dependency-injection/9781789132762/90b7cf03-764d-4eef-9f88-94de4b166c04.xhtml

func EncryptV2(keyValue Value, monitor Monitor, data []byte) ([]byte, error) {
	// As this operation make take too long, we need to be able to kill it
	stop := monitor.Done()
	result := make(chan []byte, 1)

	go func() {
		defer close(result)

		// pull the encryption key from Value
		keyRaw := keyValue.Value("encryption-key")
		if keyRaw == nil {
			panic("encryption key not found in context")
		}
		key := keyRaw.([]byte)

		// perform encryption
		ciperText := performEncryption(key, data)

		// signal complete by sending the result
		result <- ciperText
	}()

	select {
	case ciperText := <-result:
		// happy path
		return ciperText, nil

	case <-stop:
		// cancelled
		return nil, errors.New("operation cancelled")
	}
}

func performEncryption(key []byte, data []byte) []byte {
	return nil
}