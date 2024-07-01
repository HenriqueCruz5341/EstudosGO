package main

import (
	"crud-api/src/controller"
	"crud-api/src/model/repository"
	"crud-api/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(repo)
	return controller.NewUserController(userService)
}
