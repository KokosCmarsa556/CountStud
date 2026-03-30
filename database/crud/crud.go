package crud

import (
	structerr "CountStud/error"
	"CountStud/userst/student"
	User "CountStud/userst/user"

	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func CreateTableUser(ctx context.Context, conn *pgx.Conn) error {
	sqlCreate := `
		CREATE TABLE IF NOT EXISTS users (
		Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		Email VARCHAR(30) NOT NULL
		Password VARCHAR(30) NOT NULL
		Name VARCHAR(150) NOT NULL,
		SurName VARCHAR(150) NOT NULL,
		LastName VARCHAR(150) NOT NULL,
		Role VARCHAR(10) NOT NULL,
	`
	_, err := conn.Exec(ctx, sqlCreate)
	if err != nil {
		return err
	}
	return nil
}

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

func InsertRowUser(ctx context.Context, conn *pgx.Conn, u *User.User) error {

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

func CreateTableStudent(ctx context.Context, conn *pgx.Conn) error {
	sqlCreate := `
		CREATE TABLE IF NOT EXISTS students (
			Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			Name VARCHAR(150) NOT NULL,
			FirstName VARCHAR(150) NOT NULL,
			LastName VARCHAR(150) NOT NULL,
			Gender VARCHAR(50) NOT NULL,
			Address VARCHAR(150),
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

func DeleteRowInStudent(ctx context.Context, conn *pgx.Conn, u *student.Student) error {

	sqlDelete := `
		DELETE FROM students WHERE id = $1 RETURNING id;
	`

	cmdTag, err := conn.Exec(ctx, sqlDelete, u.Id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {

		log.Fatal("Student not found")

		return fmt.Errorf("Student not found")
	}

	return nil
}

func GetAllStudent(ctx context.Context, conn *pgx.Conn) ([]student.Student, error) {
	var students []student.Student
	sqlGetAll := `
		SELECT * FROM students
	`

	rows, err := conn.Query(ctx, sqlGetAll)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u student.Student
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

func GetStudentByID(ctx context.Context, conn *pgx.Conn, id uuid.UUID) (*student.Student, error) {

	var student student.Student

	sqlGetId := `
		SELECT id, name, firstName, lastName, gender, address, iin 
    FROM students 
    WHERE id = $1;
	`

	row := conn.QueryRow(ctx, sqlGetId, id)

	if err := row.Scan(
		&student.Id,
		&student.Name,
		&student.FirstName,
		&student.LastName,
		&student.Gender,
		&student.Address,
		&student.IIN); err != nil {

		if err == pgx.ErrNoRows {

			return nil, fmt.Errorf("student not found")
		}

		return nil, err
	}

	return &student, nil
}

func InsertRowInStudnet(ctx context.Context, conn *pgx.Conn, student *student.Student) error {

	sqlInsert := `
		INSERT INTO students (Name, FirstName, LastName, Gender, Address, Iin)
		VALUES($1, $2, $3, $4, $5, $6);
	`

	if _, err := conn.Exec(
		ctx,
		sqlInsert,
		student.Name,
		student.FirstName,
		student.LastName,
		student.Gender,
		student.Address,
		student.IIN); err != nil {
		return err
	}

	return nil
}

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
