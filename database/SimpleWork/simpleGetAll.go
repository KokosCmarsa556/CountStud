package simpleWork

import (
	user "CountStud/student"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllStudent(ctx context.Context, conn *pgx.Conn) ([]user.User, error) {
	var students []user.User
	sqlGetAll := `
		SELECT * FROM students
	`

	rows, err := conn.Query(ctx, sqlGetAll)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u user.User
		if err := rows.Scan(
			&u.Id,
			&u.Name,
			&u.FirstName,
			&u.LastName,
			&u.Gender,
			&u.Address,
			&u.IIN,
		); err != nil {
			return nil, err
		}
		students = append(students, u)
	}
	defer rows.Close()
	return students, rows.Err()
}
