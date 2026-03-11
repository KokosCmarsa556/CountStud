package main

import (
	user "CountStud/User"
	simpleWork "CountStud/database/SimpleWork"
	"CountStud/database/connection"
	"CountStud/handlers"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	student := user.NewUser()

	ginRoute := gin.Default()
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
	httpHandler := handlers.NewHttpHandlers(student, conn)
	// Выполняем создание таблицы
	if err := simpleWork.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
	}

	ginRoute.POST("/student", httpHandler.HandleCreateStudent)
	// Вставляем в БД
	// if err := simplework.InsertRow(ctx, conn, user.User.Name); err != nil {
	// 	log.Fatal(err)
	// }
	ginRoute.Run(":8989")
}
