package simpleWork

import (
	structerr "CountStud/structerr"
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlCreate := `
		CREATE TABLE students (
			Id SERIAL PRIMARY KEY,
			Name VARCHAR(150) NOT NULL,
			First_Name VARCHAR(150) NOT NULL,
			Last_Name VARCHAR(150) NOT NULL,
			Gender VARCHAR(50) NOT NULL,
			Address VARCHAR(150),
			Iin INTEGER NOT NULL
		)
	`

	_, err := conn.Exec(ctx, sqlCreate)
	if err != nil {
		errSt := structerr.NewErr(err.Error())
		return errSt
	}

	return nil
}
