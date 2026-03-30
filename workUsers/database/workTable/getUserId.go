package worktable

import (
	User "CountStud/workUsers/users"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetUser(ctx context.Context, conn *pgx.Conn, email string) (User.User, error) {
	var user User.User
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
			return User.User{}, fmt.Errorf("user not found")
		}
		return User.User{}, err
	}
	return user, nil
}
