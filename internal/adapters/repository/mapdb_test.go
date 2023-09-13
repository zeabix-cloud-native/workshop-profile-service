package repository

import (
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/domain"

	"testing"
)

func TestMapDBSaveNewProfile(t *testing.T) {
	repo := NewMapDBRepository()
	profile := new(domain.UserProfile)
	profile.ID = "test-01"
	profile.Username = "unittest"
	profile.Firstname = "Unit"
	profile.Lastname = "Test"
	profile.Avatar = "https://placeholder.com/unittest.png"

	err := repo.Save(profile)
	if err != nil {
		t.Errorf("Save profile return error, expected nil")
	}
}

func TestMapDBSaveExistingProfile(t *testing.T) {
	repo := NewMapDBRepository()
	profile := new(domain.UserProfile)
	profile.ID = "test-01"
	profile.Username = "unittest"
	profile.Firstname = "Unit"
	profile.Lastname = "Test"
	profile.Avatar = "https://placeholder.com/unittest.png"

	err := repo.Save(profile)
	if err != nil {
		t.Errorf("Save profile return error, expected nil")
	}

	update := new(domain.UserProfile)
	update.ID = "test-01"
	update.Username = "updated"
	update.Firstname = "Update"
	update.Lastname = "Update"
	update.Avatar = "https://placeholder.com/updated.png"

	err = repo.Save(update)
	if err != nil {
		t.Errorf("Save profile return error, expected nil")
	}

}

func TestMapDBGetProfile(t *testing.T) {
	repo := NewMapDBRepository()
	profile := new(domain.UserProfile)
	profile.ID = "test-01"
	profile.Username = "unittest"
	profile.Firstname = "Unit"
	profile.Lastname = "Test"
	profile.Avatar = "https://placeholder.com/unittest.png"

	repo.Save(profile)

	// Test Get profile
	p, err := repo.GetProfileByID("test-01")
	if err != nil {
		t.Errorf("GetProfileByID return error, expected nil")
	}

	if p.ID != "test-01" {
		t.Errorf("GetProfileByID return wrong profile, id: %s, expected test-01", p.ID)
	}
}

func TestMapDBGetProfileNotFound(t *testing.T) {
	repo := NewMapDBRepository()
	profile := new(domain.UserProfile)
	profile.ID = "test-01"
	profile.Username = "unittest"
	profile.Firstname = "Unit"
	profile.Lastname = "Test"
	profile.Avatar = "https://placeholder.com/unittest.png"

	repo.Save(profile)

	// Test Get profile
	p, err := repo.GetProfileByID("test-02")
	if err != ErrProfileNotFound {
		t.Errorf("GetProfileByID with incorrect id return %v, expected %v", err, ErrProfileNotFound)
	}

	if p != nil {
		t.Errorf("GetProfileByID with incorrect id return %v, expected nil", p)
	}
}
