package simplework

import (
	user "CountStud/User"
	"context"
	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, u *user.User) error {
	sqlDelete := `
		DELETE FROM students WHERE id = $1
	`

	if _, err := conn.Exec(ctx, sqlDelete, u.Id); err != nil {
		return err
	}
	return nil
}