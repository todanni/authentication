package repository

import "github.com/todanni/authentication/pkg/account"

func (r *repo) InsertAuthDetails(details account.AuthDetails) (account.AuthDetails, error) {
	panic("implement me")
}

func (r *repo) UpdateAuthDetails(details account.AuthDetails) (account.AuthDetails, error) {
	panic("implement me")
}

func (r *repo) GetAuthDetails(email string) (account.AuthDetails, error) {
	var ad account.AuthDetails
	err := r.db.Where(&account.AuthDetails{Email: email}).First(&ad).Error
	return ad, err
}
