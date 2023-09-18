package services

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/ports"

	"github.com/google/uuid"
)

type service struct {
	repo ports.ProfileRepository
}

func NewProfileService(r ports.ProfileRepository) ports.ProfileService {
	return &service{
		repo: r,
	}
}

func (s *service) CreateProfile(profile *domain.UserProfile) error {
	// User Profile ID will be auto generated using UUID
	profile.ID = uuid.New().String()

	return s.repo.Save(profile)
}

func (s *service) GetProfile(id string) (*domain.UserProfile, error) {
	return s.repo.GetProfileByID(id)
}

func (s *service) GetProfileByOID(oid string) (*domain.UserProfile, error) {
	return s.repo.GetProfileByOID(oid)
}
