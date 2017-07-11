package model

type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}
