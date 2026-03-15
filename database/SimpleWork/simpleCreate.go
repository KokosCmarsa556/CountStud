package simpleWork

import (
	structerr "CountStud/structerr"
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlCreate := `
		CREATE TABLE IF NOT EXISTS students (
			Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			FirstName VARCHAR(150) NOT NULL, -- Вася
			LastName VARCHAR(150) NOT NULL, -- Пупкин
			SurName VARCHAR(150) NOT NULL, -- Александрович
			Gender VARCHAR(50) NOT NULL,
			Address VARCHAR(150),
			Role TEXT, -- teacher/student
			Iin VARCHAR(12) NOT NULL
		)
	`

	_, err := conn.Exec(ctx, sqlCreate)
	if err != nil {
		errSt := structerr.NewErr(err.Error())
		return errSt
	}

	return nil
}
