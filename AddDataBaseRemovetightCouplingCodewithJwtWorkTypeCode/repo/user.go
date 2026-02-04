package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
}


type UserRepo interface {

	Creat(user User) (*User,error)

	Find(email,pass string) (*User,error)

	// List() ([]*User,error)

	// Delete(userId int) error

	// Update(user User) (*User,error)
}

type userRepo struct {
	 db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}


/* CREATE USER */
func (r userRepo) Creat(user User) (*User, error) {

	query := `
		INSERT INTO users (
		 first_name,
		 last_name,
		 email, 
		 password,
		is_shop_owner
		)
		VALUES (
		 :first_name,
		 :last_name,
		 :email, 
		 :password,
		 :is_shop_owner
		)
		RETURNING id
	`

	var userID int 

	rows, err := r.db.NamedQuery(query,user)

	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	if rows.Next(){
		rows.Scan(&userID)
	}

	user.ID = userID

	return &user,nil
}

/* FIND USER (LOGIN) */
func (r *userRepo) Find(email, pass string) (*User, error) {

	var user User

	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
	`

	err := r.db.Get(&user, query, email, pass)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
		return nil, err
	}

	return &user, nil
}