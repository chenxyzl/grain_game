package table

type colLoader[T any] struct {
	data T
}

// Load  unmarshal from csv
func (l *colLoader[T]) load(file string) error {
	return nil
}

// AfterLoad overwrite for data
func (l *colLoader[T]) afterLoad() error {
	return nil
}

// Check overwrite for check
func (l *colLoader[T]) check() error {
	return nil
}
