package ports

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/domain"

	"errors"
)

var (
	ErrProfileNotFound = errors.New("profile not found")
)

type ProfileService interface {
	CreateProfile(profile *domain.UserProfile) error
	GetProfile(id string) (*domain.UserProfile, error)
	GetProfileByOID(oid string) (*domain.UserProfile, error)
}

type ProfileRepository interface {
	Save(profile *domain.UserProfile) error
	GetProfileByID(id string) (*domain.UserProfile, error)
	GetProfileByOID(oid string) (*domain.UserProfile, error)
}
