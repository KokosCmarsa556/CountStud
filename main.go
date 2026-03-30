package main

import (
	"CountStud/workStudent/database/connection"
	simpleWork "CountStud/workStudent/database/simpleWork"
	"CountStud/workStudent/handlers"
	"context"
	"fmt"
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

	conn, err := connection.CreateConnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	httpHandler := handlers.NewHttpHandlers(conn)

	if err := simpleWork.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
	}

	//добавить обработчик обновления данных студента
	ginRoute.POST("/user/authorization", httpHandler.HandlerEntrance)
	ginRoute.POST("/user/registration", httpHandler.HandlerCreateUser)
	ginRoute.POST("/student/createstudent", httpHandler.HandlerCreateStudent)
	ginRoute.GET("/student", httpHandler.HandlerGetAllStudents)
	ginRoute.PATCH("/student/:id", httpHandler.HandlerPatchStudent)
	ginRoute.GET("/student/:id", httpHandler.HandlerGetStudentID)
	ginRoute.DELETE("/student/:id", httpHandler.HandlerDeleteStudent)

	fmt.Println("Сервер функционирует. Сервер работает на порту 8989")
	ginRoute.Run(":8989")

}
