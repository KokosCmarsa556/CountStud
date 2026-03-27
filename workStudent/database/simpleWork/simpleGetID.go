package simpleWork

import (
	user "CountStud/workStudent/student"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func GetStudentByID(ctx context.Context, conn *pgx.Conn, id uuid.UUID) (*user.Student, error) {
	//Создаем пустого студента куда в дальнейшем будут записываться данные пользователя, которого ищут
	var student user.Student

	// Создание SQL-запроса для получения пользователя по его id
	sqlGetId := `
		SELECT id, name, firstName, lastName, gender, address, iin 
    FROM students 
    WHERE id = $1;
	`
	/*
		1) Отправка SQL-запроса в PostgreSQL
		2) Передается ID пользователя котого мы ищем
		3) Возвращется строка с данными о пользователи
	*/
	row := conn.QueryRow(ctx, sqlGetId, id)

	//Scan-читает данные из row и записывает их в переменные в которых они должны хранится
	if err := row.Scan(
		&student.Id,
		&student.Name,
		&student.FirstName,
		&student.LastName,
		&student.Gender,
		&student.Address,
		&student.IIN); err != nil {
		//Проверка на то, что студент был не найден
		if err == pgx.ErrNoRows {
			//Возвращаем пустое значение и ошибку
			return nil, fmt.Errorf("student not found")
		}
		//Возвращаем пустое значение и ошибку
		return nil, err
	}
	//Возвращаем пользователя и отсутствие ошибки
	return &student, nil
}
