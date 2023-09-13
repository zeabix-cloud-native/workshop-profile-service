package main

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/handlers/httpv1"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/adapters/repository"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/services"
)

func main() {

	repo := repository.NewMapDBRepository()
	s := services.NewProfileService(repo)

	handler := httpv1.NewHttpHandlerV1(s)
	handler.Serve(":3000")
}
