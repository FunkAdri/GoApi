package models

import (
	"log"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

// Ajoute un nouvel utilisateur
func NewUser(u *User) {
	if u == nil {
		log.Fatal(u)
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO users (username, password, mail, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;", u.Username, u.Mail, u.Password, u.CreatedAt, u.UpdatedAt).Scan(&u.ID)

	if err != nil {
		log.Fatal(err)
	}
}

// Récupère un seul utilisateur
func FindUserById(id int) *User {
	var user User

	row := config.Db().QueryRow("SELECT * FROM users WHERE id = $1;", id)
	err := row.Scan(&user.ID, &user.Username, &user.Mail, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &user
}

// Récupère tous les utilisateurs
func AllUsers() *Users {
	var users Users

	rows, err := config.Db().Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Username, &u.Mail, &u.CreatedAt, &u.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}

	return &users
}

func UpdateUser(user *User) {
	user.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE users SET username=$2, mail=$3, password=$4, updated_at=$5 WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Username, user.Mail, user.Password, user.UpdatedAt, user.ID)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUserById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM users WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
