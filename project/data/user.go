package data

import (
	"database/sql"
	"errors"
)

type User struct {
	UserName string
	Email    string
	Age      int32
}

type UserModel struct {
	DB *sql.DB
}

func (um *UserModel) AddUser(username, email string, age int32) error {
	query := `
	INSERT INTO users(username, email, age)
	VALUES($1, $2, $3)
	`
	_, err := um.DB.Exec(query, username, email, age)
	if err != nil {
		return errors.New("Couldn't add user to the db")
	}
	return nil
}

func (um *UserModel) CheckUser(userID int32) (*User, error) {
	query := `
	SELECT username, email, age
	FROM users
	WHERE id = $1
	`
	var user User
	err := um.DB.QueryRow(query, userID).Scan(&user.UserName, &user.Email, &user.Age)
	if err != nil {
		return nil, errors.New("Couldn't fetch user")
	}
	return &user, nil
}
