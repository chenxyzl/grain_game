package table

type rowLoader[k comparable, T any] struct {
	list []T
	data map[k]T
}

// Load  unmarshal from csv
func (l *rowLoader[K, T]) load(file string) error {
	return nil
}

// AfterLoad overwrite for data
func (l *rowLoader[K, T]) afterLoad() error {
	return nil
}

// Check overwrite for check
func (l *rowLoader[K, T]) check() error {
	return nil
}

func (l *rowLoader[K, T]) GetById(id K) T {
	return l.data[id]
}

func (l *rowLoader[K, T]) Range(f func(k K, v T) bool) {
	for k, v := range l.data {
		f(k, v)
	}
}
