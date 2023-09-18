package main

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/handlers"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/repository"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/services"
	"github.com/zeabix-cloud-native/workshop-profile-service/utils"
)

func main() {

	repo := repository.NewMapDBRepository()
	// Init DB
	utils.InitDB()

	s := services.NewProfileService(repo)

	handler := handlers.NewHTTPHandler(s)
	handler.Serve("3001")
}
