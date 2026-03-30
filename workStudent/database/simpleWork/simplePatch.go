package simpleWork

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// type studentDTO struct {
// 	student *student.Student
// }

func PatchStudent(ctx context.Context, conn *pgx.Conn, id uuid.UUID, name, lastname, address string) error {

	if name != "" && lastname != "" && address != "" {
		sqlPathUser := `
		UPDATE students
		SET name = $1, lastname = $2, Address = $3
		WHERE id = $4
	`
		if _, err := conn.Exec(ctx, sqlPathUser, name, lastname, address); err != nil {
			return err
		}
	}
	if name == "" && lastname == "" && address == "" {
		return nil
	}
	if name != "" {
		sqlPathUser := `
		UPDATE students
		SET name = $1, 
		WHERE id = $2
	`
		if _, err := conn.Exec(ctx, sqlPathUser, name); err != nil {
			return err
		}
	}

	if lastname != "" {
		sqlPathUser := `
		UPDATE students
		SET lastname = $1, 
		WHERE id = $2
	`
		if _, err := conn.Exec(ctx, sqlPathUser, lastname); err != nil {
			return err
		}
		return nil
	}

	if address != "" {
		sqlPathUser := `
		UPDATE students
		SET address = $1, 
		WHERE id = $2
	`
		if _, err := conn.Exec(ctx, sqlPathUser, address); err != nil {
			return err
		}
		return nil
	}

	return nil
}
