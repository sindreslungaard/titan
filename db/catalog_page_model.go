package db

type CatalogPage struct {
	ID       int
	ParentID int
	Order    int
	Name     string
	Icon     int
	MinRank  int
	Layout   string
	Text     string
}
