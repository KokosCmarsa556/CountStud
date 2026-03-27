package worktable

import (
	"CountStud/workUsers/users"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetUser(ctx context.Context, conn *pgx.Conn, email string) (users.Users, error) {
	var user users.Users
	sqlGetUserId := `
		SELECT Id, Email, Password, Name, SurName, LastName, Role 
    FROM students 
    WHERE Email = $1;
	`

	row := conn.QueryRow(ctx, sqlGetUserId, email)
	if err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.SurName,
		&user.LastName,
		&user.Role,
	); err != nil {
		if err == pgx.ErrNoRows {
			return users.Users{}, fmt.Errorf("user not found")
		}
		return users.Users{}, err
	}
	return user, nil
}
