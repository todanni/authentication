package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/todanni/authentication/pkg/account"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) InsertAccount(acc account.Account) (account.Account, error) {
	args := r.Called(acc)
	return args.Get(0).(account.Account), args.Error(1)
}

func (r *RepositoryMock) GetAuthDetails(email string) (account.AuthDetails, error) {
	args := r.Called(email)
	return args.Get(0).(account.AuthDetails), args.Error(1)
}

func (r *RepositoryMock) InsertVerificationRecord(record account.VerificationRecord) (account.VerificationRecord, error) {
	args := r.Called(record)
	return args.Get(0).(account.VerificationRecord), args.Error(1)
}

func (r *RepositoryMock) GetVerificationRecordByCode(code string) (account.VerificationRecord, error) {
	args := r.Called(code)
	return args.Get(0).(account.VerificationRecord), args.Error(1)
}

func (r *RepositoryMock) SetAuthDetailsValid(accountID uint) error {
	args := r.Called(accountID)
	return args.Error(0)
}
