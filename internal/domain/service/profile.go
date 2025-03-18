package service

import (
	"context"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
)

type ProfileService interface {
	GetProfile(ctx context.Context, userID uint) (*domain.Profile, error)
	UpdateProfile(ctx context.Context, profile *domain.Profile) error
}

type profileService struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(profileRepo repository.ProfileRepository) ProfileService {
	return &profileService{
		profileRepo: profileRepo,
	}
}

func (s *profileService) GetProfile(ctx context.Context, userID uint) (*domain.Profile, error) {
	profile, err := s.profileRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, errors.CodeNotFound, "用户资料不存在")
	}
	return profile, nil
}

func (s *profileService) UpdateProfile(ctx context.Context, profile *domain.Profile) error {
	if err := s.profileRepo.Update(ctx, profile); err != nil {
		return errors.Wrap(err, errors.CodeServerError, "更新用户资料失败")
	}
	return nil
}