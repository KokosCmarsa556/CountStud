package main

import (
	"CountStud/database/connection"
	"CountStud/database/simplework"
	"CountStud/handlers"
	"CountStud/user"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	student := user.NewUser()

	ginRoute := gin.Default()

	ginRoute.ServeHTTP(":8989")
	httpHandler := handlers.NewHttpHandlers(student)

	ctx := context.Background()
	//Accessing the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Создаем подключение к базе данных
	conn, err := connection.CreateConnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Выполняем создание таблицы
	if err := simplework.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
	}

	// Вставляем в БД
	// if err := simplework.InsertRow(ctx, conn, user.User.Name); err != nil {
	// 	log.Fatal(err)
	// }
}
