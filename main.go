package main

import (
	"CountStud/workStudent/database/connection"
	simpleWork "CountStud/workStudent/database/simpleWork"
	"CountStud/workStudent/handlers"
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
	ginRoute.GET("/student/:id", httpHandler.HandlerGetStudentID)
	ginRoute.GET("/students", httpHandler.HandlerGetAllStudents)
	ginRoute.DELETE("/student/:id", httpHandler.HandlerDeleteStudent)
	// Вставляем в БД
	// if err := simplework.InsertRow(ctx, conn, user.User.Name); err != nil {
	// 	log.Fatal(err)
	// }
	ginRoute.Run(":8989")
}
