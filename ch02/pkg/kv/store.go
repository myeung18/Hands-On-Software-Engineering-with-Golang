package kv

type Store struct {
	name string
}

func (s Store) Put(key string, value interface{}) error {
	return nil
}

func (s Store) Close() error {

	return nil
}
