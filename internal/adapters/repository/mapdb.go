package repository

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/ports"
)

type mapdb struct {
	DB map[string]*domain.UserProfile
}

func NewMapDBRepository() ports.ProfileRepository {
	return &mapdb{
		DB: make(map[string]*domain.UserProfile),
	}
}

func (repo *mapdb) Save(profile *domain.UserProfile) error {
	// Check if there is existing Profile
	p, ok := repo.DB[profile.ID]
	if ok {
		// Update existing profile
		p.Username = profile.Username
		p.Firstname = profile.Firstname
		p.Lastname = profile.Lastname
		p.Avatar = profile.Avatar
	} else {
		repo.DB[profile.ID] = profile
	}

	return nil
}

func (repo *mapdb) GetProfileByID(id string) (*domain.UserProfile, error) {
	p, ok := repo.DB[id]
	if !ok {
		return nil, ports.ErrProfileNotFound
	}

	return p, nil
}

func (repo *mapdb) GetProfileByOID(oid string) (*domain.UserProfile, error) {
	for _, v := range repo.DB {
		if v.OID == oid {
			return v, nil
		}
	}

	return nil, ports.ErrProfileNotFound
}
