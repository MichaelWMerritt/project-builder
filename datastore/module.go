package datastore

type Module struct {
	Collection
}

func NewModuleDatastore() Module {
	return Module{Collection{}}
}