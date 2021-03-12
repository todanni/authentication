package repository

import "github.com/todanni/authentication/pkg/account"

func (r repo) InsertAuthDetails(details account.AuthDetails) (account.AuthDetails, error) {
	panic("implement me")
}

func (r repo) UpdateAuthDetails(details account.AuthDetails) (account.AuthDetails, error) {
	panic("implement me")
}

func (r repo) GetAuthDetails(userID int) (account.AuthDetails, error) {
	panic("implement me")
}
