package model

type DaoModel interface {
	GetPrimaryKey() int
}

type DtoModel interface {
	String() string
}
