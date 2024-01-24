package entity

type Ingredient struct {
	Id       uint
	Name     string
	Quantity int64
	Unity    *string
	Types    Type
}
