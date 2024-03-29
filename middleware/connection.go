package middleware

import (
	//"bufio"
	"context"
	//"fmt"
	"os"
	//"strings"
	"log"

	"github.com/jsatof/qbook/models"

	"github.com/jackc/pgx/v5"
)

type User = models.User

var db *pgx.Conn

func GetUserWithMaxID() (User, error) {
	rows, err := db.Query(context.Background(), "SELECT * FROM users ORDER BY id DESC LIMIT 1")
	user := User{}
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(user)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func InsertUser(u User) error {
	return nil
}

func InitDB() error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func Close() {
	err := db.Close(context.Background())
	if err != nil {
		log.Fatalf("Failed to close the database connection:\n%s\n", err)
	}
	log.Println("Closing database connection")
}
