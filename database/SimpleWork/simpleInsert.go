package simplework

import (
	user "CountStud/User"
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, student *user.User) error {
	sqlInsert := `
		INSERT INTO students (Name, First_Name, Last_Name, Gender, Address, Iin)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	if _, err := conn.Exec(ctx, sqlInsert, student.Name, student.FirstName, student.LastName, student.Gender, student.Address, student.IIN); err != nil {
		return err
	}
	return nil
}
