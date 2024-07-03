package table

// iLoader 加载接口的实现
type iLoader interface {
	load(file string) error
	afterLoad() error
	check() error
}

// finalLoader iLoader之后的实现,一般用于在iLoad加载完成后组装数据重新构成表
type finalLoader interface {
	load() error
	afterLoad() error
	check() error
}
