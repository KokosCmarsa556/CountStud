package simpleWork

import (
	user "CountStud/workStudent/student"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllStudent(ctx context.Context, conn *pgx.Conn) ([]user.Student, error) {
	var students []user.Student
	sqlGetAll := `
		SELECT * FROM students
	`

	rows, err := conn.Query(ctx, sqlGetAll)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u user.Student
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
	return students, rows.Err()
}
