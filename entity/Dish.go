package entity

type Dish struct {
	Id          uint
	Name        string
	Description string
	Steps       []Step
	Ingredients []Ingredient
}
