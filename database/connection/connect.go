package connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConnect(ctx context.Context) (*pgx.Conn, error) {

	//protocol://username:password@host:port/database
	dbURL := os.Getenv("DB_URL")

	return pgx.Connect(ctx, dbURL)
}
