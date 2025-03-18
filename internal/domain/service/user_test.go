package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestUserService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	ctx := context.Background()
	username := "testuser"
	password := "password123"
	email := "test@example.com"

	mockRepo.On("GetByUsername", ctx, username).Return(nil, nil)
	mockRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).Return(nil)

	err := service.Register(ctx, username, password, email)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}