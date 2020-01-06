package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialise la base de donnée
func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=magiclibrary_admin dbname=magiclibrary")

	if err != nil {
		log.Fatal(err)
	}

	// Create Table users if not exists
	createUsersTable()
}

// Créer la table des utilisateurs
func createUsersTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY NOT NULL, username varchar(20), mail varchar(255), password varchar(255), created_at timestamp default NULL, updated_at timestamp default NULL)")

	if err != nil {
		log.Fatal(err)
	}
}

// Accès en lecture pour db var
func Db() *sql.DB {
	return db
}
