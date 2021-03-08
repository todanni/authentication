package repository

import (
	"context"
	_ "github.com/lib/pq" // here
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/todanni/authentication/pkg/auth"
	"github.com/todanni/authentication/test/container"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

type AuthenticationTestSuite struct {
	suite.Suite
	db        *gorm.DB
	container *container.PgContainer
}

func TestRunAuthenticationTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}

// SetupTest runs before each test.
func (suite *AuthenticationTestSuite) SetupTest() {
	suite.cleanupDatabase()
}

func (suite *AuthenticationTestSuite) TestInsert() {
	repo := NewAuthRepository(suite.db)

	authDetails := auth.AuthenticationDetails{
		Email:    "test@email.com",
		Password: "Password123",
	}

	created, err := repo.Insert(authDetails)
	assert.NoError(suite.T(), err)

	// Verify values
	assert.Equal(suite.T(), created.Email, "test@email.com")
	assert.Equal(suite.T(), created.Password, "Password123")
	assert.Equal(suite.T(), created.Verified, false)
	assert.Equal(suite.T(), created.AccountID, uint(0))
}

func (suite *AuthenticationTestSuite) SetupSuite() {
	pgContainer, err := container.NewPGContainer("database_for_it")
	if err != nil {
		suite.T().Fatalf("couldn't start container: %v", err)
	}

	suite.container = pgContainer

	connection, err := pgContainer.ConnectionString()
	assert.NoError(suite.T(), err)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Do migrations
	err = db.AutoMigrate(&auth.AuthenticationDetails{})
	assert.NoError(suite.T(), err)

	suite.db = db

	suite.cleanupDatabase()
}

func (suite *AuthenticationTestSuite) TearDownSuite() {
	suite.cleanupDatabase()
	assert.NoError(suite.T(), suite.container.Close(context.Background()))
}

func (suite *AuthenticationTestSuite) cleanupDatabase() {
	err := suite.db.Exec("DELETE FROM authentication_details")
	assert.NoError(suite.T(), err.Error)
}
