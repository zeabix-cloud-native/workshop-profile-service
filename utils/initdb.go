package utils

import (
	"log"

	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/ports"
)

func InitDB(repo ports.ProfileRepository) {
	log.Println("TODO")

	p := new(domain.UserProfile)
	p.ID = "4022a817-365c-407c-880b-c46ce6cf8880"
	p.Username = "test"
	p.Firstname = "test"
	p.Lastname = "test"
	p.Avatar = ""
	p.DOB = "1/1/2022"
	p.Mobile = "08898767483"
	p.Address = "Anywhere"
	p.OID = "3fb0150a-267c-43f3-a86e-c2a9fd495500"

	repo.Save(p)
}
