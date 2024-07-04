package table

type colLoader[T any] struct {
	data T
}

// load  unmarshal from csv
func (l *colLoader[T]) load(file string) error {
	//todo parse excel
	return nil
}

// afterLoad overwrite for data
func (l *colLoader[T]) afterLoad() error {
	return nil
}

// check overwrite for check
func (l *colLoader[T]) check() error {
	return nil
}

// Get return data
func (l *colLoader[T]) Get() T {
	return l.data
}
