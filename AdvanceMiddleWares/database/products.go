package database


type Products struct {
    
	ID int               `json:"id"`
	Title string         `json:"title"`
	Description string   `json:"description"`
	Price float64        `json:"price"`
	ImgUrl string        `json:"imgUrl"`
}

var ProductsList []Products


func init() {
	prd1 := Products{
		ID:          1,
		Title:       "Mango",
		Description: " This is my favorite food",
		Price:       205,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSuQfOwRMG2Fk_6D880IWnkQf6BfkGEKjN7oQ&s",
	}
	prd2 := Products{
		ID:          2,
		Title:       "Banana",
		Description: " This is my second favorite food",
		Price:       20,
		ImgUrl:      "https://img.freepik.com/free-photo/bananas-white-background_1187-1671.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prd3 := Products{
		ID:          3,
		Title:       "Watermelon",
		Description: " This food is very testy",
		Price:       580,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRKZq-G705ZlMOJh69vQ4p3oEtDWogurMz2sQ&s",
	}
	prd4 := Products{
		ID:          4,
		Title:       "Jackfruit",
		Description: " i do not like this food because is eating boring",
		Price:       930,
		ImgUrl:      "https://img.freepik.com/free-photo/yellow-jackfruit_74190-4803.jpg?semt=ais_hybrid&w=740&q=80",
	}
	prd5 := Products{
		ID:          5,
		Title:       "Lichi",
		Description: " I like this food because it is my fevorite food like ever",
		Price:       360,
		ImgUrl:      "https://www.kisanshop.in/uploads/muzaffarpur-lichi-plant.webp",
	}

	prd6 := Products{
		ID:          6,
		Title:       "durian jackfruit",
		Description: "This is a durian Fruit . It is various is jackfruit",
		Price:       2652,
		ImgUrl:      "https://howdaily.com/wp-content/uploads/2019/09/durians-rinds-800x533.jpg?x19738",
	}

	prd7 := Products{
		ID:          7,
		Title:       "green coconut",
		Description: "This is a Green Cocont . it called in Daab",
		Price:       150,
		ImgUrl:      "https://www.sampanresort.com/images/fresh-green-coconut-1.jpg",
	}

	ProductsList = append(ProductsList, prd1)
	ProductsList = append(ProductsList, prd2)
	ProductsList = append(ProductsList, prd3)
	ProductsList = append(ProductsList, prd4)
	ProductsList = append(ProductsList, prd5)
	ProductsList = append(ProductsList, prd6)
	ProductsList = append(ProductsList, prd7)
}