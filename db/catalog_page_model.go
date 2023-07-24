package db

type CatalogPage struct {
	ID       int
	ParentID int
	OrderNum int
	Name     string
	Icon     int
	Layout   string
	Text     string
}
