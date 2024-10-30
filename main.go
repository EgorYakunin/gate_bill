package main

import (
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
    Id          int       `json:"id"`
    Email       string    `json:"email"`
    Username    string    `json:"username"`
    CreatedAt   time.Time `json:"created_at"`
    TelegramId  string    `json:"telegram_id"`
    AvatarPath  string    `json:"avatar_path"`
    FamilyId    int       `json:"family_id"`
}

func dbConnect() *sqlx.DB {
	connStr := "user=egor password=password123 dbname=gate_bill_db sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		panic("db: can't open sql connection")
	}

	err = db.Ping()
	if err != nil {
		panic("db: can't ping a database")
	}

	return db
}

func createUser(db *sqlx.DB, user *User) (*User, error) {
	query := `insert into app_user (email, username) values (:email, :username) returning id`

	var id int
	rows, err := db.NamedQuery(query, user)
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			panic("Scan failed")	
		}
	}

	if err != nil {
		fmt.Println(err)	
		panic("fuuuuuuck")
	}
	
	user.Id = id
	return user, nil
}

func main() {
	db := dbConnect()	
	defer db.Close()
	
	newUser := &User{
		Username: "Andrew",
		Email: "andrew@gmail.com",
	}

	user, err := createUser(db, newUser)
	if err != nil {
		fmt.Println(err)
		panic("on noooo")
	}

	fmt.Println("Created new user id:", user.Id)
}


