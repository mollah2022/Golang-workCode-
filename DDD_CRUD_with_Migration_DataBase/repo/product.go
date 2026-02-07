package repo

import (
	"database/sql"
	"ecomerce/domain"
	"ecomerce/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}


func (r *productRepo) Creat(p domain.Products) (*domain.Products,error) {

	query := `
	INSERT INTO products (
     title,
     description,
     price,
     img_url
    )
    VALUES (
     $1,
     $2,
     $3,
     $4
    )
    RETURNING id
  `
 row := r.db.QueryRow(query,p.Title,p.Description,p.Price,p.ImgUrl)
  err := row.Scan(&p.ID)

  if err !=nil {
	return nil,err
  }

  return &p,nil

}
func (r *productRepo) Get(id int) (*domain.Products,error) {

    var prd domain.Products

	query := `
	   SELECT
	      id,
	      title,
          description,
          price,
          img_url
		from products
		where id = $1
	`
	err := r.db.Get(&prd, query, id)

	if err!=nil {
       if err == sql.ErrNoRows {
		  return nil,nil
	   }
	   return nil,err
	}

	return &prd,nil

}
func (r *productRepo) List() ([]*domain.Products,error) {

	 var prdList []*domain.Products

	query := `
	   SELECT
	      id,
	      title,
          description,
          price,
          img_url
		from products
	`
	err := r.db.Select(&prdList, query)

	if err!=nil {
	   return nil,err
	}

	return prdList,nil

}
func (r *productRepo) Delete(id int) error {

	query := `
	   DELETE FROM products WHERE id = $1
	`
	_,err := r.db.Exec(query,id)

	if err!=nil {
		return err
	}
	return nil
}
func (r *productRepo) Update(p domain.Products) (*domain.Products,error) {

	query := `
	  UPDATE products
	  SET 
	   title=$1,
	   description=$2,
	   price=$3,
	   img_url=$4
	   WHERE id = $5
	`
	row := r.db.QueryRow(query,p.Title,p.Description,p.Price,p.ImgUrl)
	err := row.Err()

	if err!=nil {
		return nil,err
	}

	return &p,nil
	
}

// func generateInitialProducts(r *productRepo) {
// 	prd1 := &Products{
// 		ID:          1,
// 		Title:       "Mango",
// 		Description: " This is my favorite food",
// 		Price:       205,
// 		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSuQfOwRMG2Fk_6D880IWnkQf6BfkGEKjN7oQ&s",
// 	}
// 	prd2 := &Products{
// 		ID:          2,
// 		Title:       "Banana",
// 		Description: " This is my second favorite food",
// 		Price:       20,
// 		ImgUrl:      "https://img.freepik.com/free-photo/bananas-white-background_1187-1671.jpg?semt=ais_hybrid&w=740&q=80",
// 	}
// 	prd3 := &Products{
// 		ID:          3,
// 		Title:       "Watermelon",
// 		Description: " This food is very testy",
// 		Price:       580,
// 		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRKZq-G705ZlMOJh69vQ4p3oEtDWogurMz2sQ&s",
// 	}
// 	prd4 := &Products{
// 		ID:          4,
// 		Title:       "Jackfruit",
// 		Description: " i do not like this food because is eating boring",
// 		Price:       930,
// 		ImgUrl:      "https://img.freepik.com/free-photo/yellow-jackfruit_74190-4803.jpg?semt=ais_hybrid&w=740&q=80",
// 	}
// 	prd5 := &Products{
// 		ID:          5,
// 		Title:       "Lichi",
// 		Description: " I like this food because it is my fevorite food like ever",
// 		Price:       360,
// 		ImgUrl:      "https://www.kisanshop.in/uploads/muzaffarpur-lichi-plant.webp",
// 	}

// 	prd6 := &Products{
// 		ID:          6,
// 		Title:       "durian jackfruit",
// 		Description: "This is a durian Fruit . It is various is jackfruit",
// 		Price:       2652,
// 		ImgUrl:      "https://howdaily.com/wp-content/uploads/2019/09/durians-rinds-800x533.jpg?x19738",
// 	}

// 	prd7 := &Products{
// 		ID:          7,
// 		Title:       "green coconut",
// 		Description: "This is a Green Cocont . it called in Daab",
// 		Price:       150,
// 		ImgUrl:      "https://www.sampanresort.com/images/fresh-green-coconut-1.jpg",
// 	}


// 	r.productsList = append(r.productsList, prd1)
// 	r.productsList = append(r.productsList, prd2)
// 	r.productsList = append(r.productsList, prd3)
// 	r.productsList = append(r.productsList, prd4)
// 	r.productsList = append(r.productsList, prd5)
// 	r.productsList = append(r.productsList, prd6)
// 	r.productsList = append(r.productsList, prd7)
// }