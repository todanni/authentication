package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq" // here
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/todanni/authentication/pkg/account"
	"github.com/todanni/authentication/test/container"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AccountRepoTestSuite struct {
	suite.Suite
	db        *gorm.DB
	container *container.PgContainer
}

func TestRunAuthenticationTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepoTestSuite))
}

// SetupTest runs before each test.
func (suite *AccountRepoTestSuite) SetupTest() {
	suite.cleanupDatabase()
}

func (suite *AccountRepoTestSuite) TestInsertAccount() {
	r := NewRepository(suite.db)

	acc := account.Account{
		FirstName:      "First",
		LastName:       "Lst",
		ProfilePicture: "http://imgur.com/happy.jpeg",
		JobTitle:       "Software Engineer",
		AuthDetails: account.AuthDetails{
			Email:    "test@mail.com",
			Password: "test",
			Verified: false,
		},
	}

	created, err := r.InsertAccount(acc)
	require.NoError(suite.T(), err)

	// Verify values
	assert.Equal(suite.T(), created.FirstName, acc.FirstName)
	assert.Equal(suite.T(), created.AuthDetails.Email, acc.AuthDetails.Email)
	assert.Equal(suite.T(), created.AuthDetails.Password, acc.AuthDetails.Password)
}

func (suite *AccountRepoTestSuite) SetupSuite() {
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
	require.NoError(suite.T(), err)

	suite.db = db
	suite.cleanupDatabase()
}

func (suite *AccountRepoTestSuite) TearDownSuite() {
	suite.cleanupDatabase()
	assert.NoError(suite.T(), suite.container.Close(context.Background()))
}

func (suite *AccountRepoTestSuite) cleanupDatabase() {
	suite.db.Exec("DELETE FROM authentication_details")
}
