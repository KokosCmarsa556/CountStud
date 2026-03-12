package main

import (
	_ "CountStud/User"
	simpleWork "CountStud/database/SimpleWork"
	"CountStud/database/connection"
	"CountStud/handlers"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
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
	httpHandler := handlers.NewHttpHandlers(conn)
	// Выполняем создание таблицы
	if err := simpleWork.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
	}

	ginRoute.POST("/student", httpHandler.HandlerCreateStudent)
	// Вставляем в БД
	// if err := simplework.InsertRow(ctx, conn, user.User.Name); err != nil {
	// 	log.Fatal(err)
	// }
	ginRoute.Run(":8989")
}
