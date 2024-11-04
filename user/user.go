package user

import (
	"fmt"
	"gate_bill/db"
	_ "github.com/lib/pq"
)

func GetUsers() ([]User, error) {
    var users []User
    err := db.Conn.Select(&users, "SELECT id, username, email FROM app_user")
    if err != nil {
        return nil, err
    }

    return users, nil
}

func CreateUser(user *User) (*User, error) {
	query := `insert into app_user (email, username) values (:email, :username) returning id`

	var id int
	rows, err := db.Conn.NamedQuery(query, user)
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

