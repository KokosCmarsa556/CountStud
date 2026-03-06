package simplework

import (
	user "CountStud/User"
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, u *user.User) error {
	sqlInsert := `
		INSERT INTO students (Name, First_Name, Last_Name, Gender, Address, Iin)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	if _, err := conn.Exec(ctx, sqlInsert, u.Name, u.FirstName, u.LastName, u.Gender, u.Address, u.IIN); err != nil {
		return err
	}
	return nil
}
