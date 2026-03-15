package connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnect(ctx context.Context) (*pgx.Conn, error) {

	// { user_id: 1 } 	

	//protocol://username:password@host:port/database
	secretKey := os.Getenv("SECRET_KEY")
	dbURL := os.Getenv("DB_URL")

	return pgx.Connect(ctx, dbURL)
}
