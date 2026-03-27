package simpleWork

import (
	user "CountStud/workStudent/student"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, conn *pgx.Conn, u *user.Student) error {
	//Создаем запрос для удаления пользователя из БД. RETURNING id - проверяет есть ли такой пользователь или нет
	sqlDelete := `
		DELETE FROM students WHERE id = $1 RETURNING id;
	`
	//Возвращем поличество удаленных строк и ошибку из запроса
	cmdTag, err := conn.Exec(ctx, sqlDelete, u.Id)
	//Проверяем на наличие ошибок
	if err != nil {
		return err
	}
	// Проверяем, что если кол-во возврщаемых строк 0, то значит, что такого пользователя нет
	if cmdTag.RowsAffected() == 0 {
		//Логируем ошибку
		log.Fatal("Student not found")
		//Возвращаем ошибку
		return fmt.Errorf("Student not found")
	}
	//Если все правильно и ошибок нет, значит пользователь удален
	return nil
}
