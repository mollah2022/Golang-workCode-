package domain

type Products struct {
	ID          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	ImgUrl      string  `db:"img_url" json:"imgUrl"`
}