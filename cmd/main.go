package main

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/handlers"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/repository"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/services"
	"github.com/zeabix-cloud-native/workshop-profile-service/utils"

	"fmt"
	"log"
)

func main() {

	port := utils.GetEnv("PORT", "8080")

	repo := repository.NewMapDBRepository()
	// Init DB
	utils.InitDB(repo)

	s := services.NewProfileService(repo)

	handler := handlers.NewHTTPHandler(s)

	log.Printf("Start profile service at port %s", port)
	handler.Serve(fmt.Sprintf(":%s", port))
}
