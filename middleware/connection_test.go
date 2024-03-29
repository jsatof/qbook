package middleware

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestDBConnection(t *testing.T) {
	db, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		t.Fatalf("Could not connect to database\n%s\n", err)
	}
	db.Close(context.Background())
}
