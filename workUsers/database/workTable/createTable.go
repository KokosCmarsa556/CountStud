package worktable

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlCreate := `
		CREATE TABLE IF NOT EXISTS users (
		Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		Login VARCHAR(30) NOT NULL
		Password VARCHAR(30) NOT NULL
		Name VARCHAR(150) NOT NULL,
		First_Name VARCHAR(150) NOT NULL,
		Last_Name VARCHAR(150) NOT NULL,
		Gender VARCHAR(50) NOT NULL,
		Role VARCHAR(10) NOT NULL,
	`
	_, err := conn.Exec(ctx, sqlCreate)
	if err != nil {
		return err
	}
	return nil
}
