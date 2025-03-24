package service

import (
	"database/sql"
	"errors"
	"testDBMock/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByID(id int) (*repository.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repository.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *repository.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestGetUser(t *testing.T) {
	// Create a new instance of our mock repository
	mockRepo := new(MockUserRepository)

	// Create a test user
	testUser := &repository.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@example.com",
	}

	// Set up expectations
	mockRepo.On("GetByID", 1).Return(testUser, nil)
	mockRepo.On("GetByID", 2).Return(nil, sql.ErrNoRows)

	// Create a new service with our mock repo
	userService := NewUserService(mockRepo)

	// Test successful user retrieval
	t.Run("Success", func(t *testing.T) {
		user, err := userService.GetUser(1)
		assert.NoError(t, err)
		assert.Equal(t, testUser, user)
	})

	// Test user not found
	t.Run("NotFound", func(t *testing.T) {
		user, err := userService.GetUser(2)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}

func TestCreateUser(t *testing.T) {
	// Create a new instance of our mock repository
	mockRepo := new(MockUserRepository)

	// Set up expectations
	mockRepo.On("Create", mock.MatchedBy(func(user *repository.User) bool {
		return user.Username == "new_user" && user.Email == "new@example.com"
	})).Return(nil)

	mockRepo.On("Create", mock.MatchedBy(func(user *repository.User) bool {
		return user.Username == "error_user"
	})).Return(errors.New("database error"))

	// Create a new service with our mock repo
	userService := NewUserService(mockRepo)

	// Test successful user creation
	t.Run("Success", func(t *testing.T) {
		err := userService.CreateUser("new_user", "new@example.com")
		assert.NoError(t, err)
	})

	// Test creation error
	t.Run("Error", func(t *testing.T) {
		err := userService.CreateUser("error_user", "error@example.com")
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
