package db

type Database[T any] interface {
	Client() interface{}
	Get(query interface{}) []T // to obtain multiple values, pass the valid object for the database instantiated
	GetById(id string) T       // get a single element, searched by ID given
	Insert(data T) error
	UpdateById(id string, data T) error
	DeleteById(id string) error
	Type() string // get the name of the database client beign used by the instance
}
