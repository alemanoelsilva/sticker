package user_test

import (
	"errors"
	"sticker/internal/app/entity"
	user "sticker/internal/app/repository/mysql/users"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type DB interface {
	Save(value interface{}) *DB
	Error() error
}

type Logger interface {
	Error() *zerolog.Event
}

// MockDB is a mock implementation of DB interface
type MockDB struct {
	mock.Mock
	*gorm.DB
}

func (m *MockDB) Save(value interface{}) *DB {
	args := m.Called(value)
	return args.Get(0).(*DB)
}

func (m *MockDB) Error() error {
	args := m.Called()
	return args.Error(0)
}

// MockLogger is a mock implementation of Logger interface
type MockLogger struct {
	mock   mock.Mock
	logger *zerolog.Logger
}

func (m *MockLogger) Error() *zerolog.Event {
	args := m.mock.Called()
	return args.Get(0).(*zerolog.Event)
}

func TestAddUserRepository(t *testing.T) {
	mockDB := &MockDB{
		Mock: ,
	}
	mockLogger := &MockLogger{}

	var err error

	userSignUp := entity.SignUp{
		Name:     "John",
		Email:    "john@example.com",
		Password: "password123",
	}

	// userRepository, _ := repository.NewSqlRepository(*mockDB, &mockLogger)

	userRepository := user.SqlRepository{DB: mockDB.DB, Logger: mockLogger.logger}

	mockDB.On("Save", mock.Anything).Return(&gorm.DB{}).Once()
	err = userRepository.AddUser(userSignUp)
	assert.NoError(t, err)

	expectedError := errors.New("database error")
	mockDB.On("Save", mock.Anything).Return(&gorm.DB{}).Once().Return(&gorm.DB{}, expectedError)
	mockLogger.mock.On("Error").Once()
	err = userRepository.AddUser(userSignUp)
	assert.EqualError(t, err, "internal Server Error")

	mockDB.AssertExpectations(t)
	mockLogger.mock.AssertExpectations(t)
}
