package simpleWork

import (
	user "CountStud/student"
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, student *user.User) error {
	//Создаем запрос. Не передаем id, т.к он генерируется у нас в PostgreSQL
	sqlInsert := `
		INSERT INTO students (FirstName, LastName, Gender, Address, Iin)
		VALUES($1, $2, $3, $4, $5, $6);
	`

	// Отправляем все данные пользователя в БД
	if _, err := conn.Exec(
		ctx,
		sqlInsert,
		student.FirstName,
		student.LastName,
		// todo
		student.Gender,
		student.Address,
		student.IIN); err != nil { // проверяем на ошибки
		return err // если ошибка есть, то возвращаем ее
	}
	// Если ошибки нет, то просто возвращаем nil
	return nil
}
