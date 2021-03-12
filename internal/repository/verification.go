package repository

import (
	"github.com/todanni/authentication/pkg/account"
)

func (r repo) InsertVerificationRecord(record account.VerificationRecord) (account.VerificationRecord, error) {
	panic("implement me")
}

func (r repo) GetVerificationRecord(accountID int) (account.VerificationRecord, error) {
	panic("implement me")
}
