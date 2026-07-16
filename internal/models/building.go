package models

type Building struct {
	Id          string `csv:"id"`
	Tile        string `csv:"title"`
	Description string `csv:"description"`
	BasePrice   int    `csv:"base_price"`
	Income      int    `csv:"income"`
}
