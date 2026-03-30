package worktable

import (
	User "CountStud/workUsers/users"
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, u *User.User) error {
	//Создаем запрос. Не передаем id, т.к он генерируется у нас в PostgreSQL
	sqlInsert := `
		INSERT INTO users (Login, Password, Name, FirstName, LastName, Gender, Role)
		VALUES($1, $2, $3, $4, $5, $6, $7);
	`

	if _, err := conn.Exec(
		ctx,
		sqlInsert,
		u.Email,
		u.Password,
		u.Name,
		u.SurName,
		u.LastName,
		u.Role); err != nil {
		return err
	}

	return nil
}
