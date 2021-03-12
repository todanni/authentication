package repository

import "github.com/todanni/authentication/pkg/account"

func (r repo) GetAccount(id int) (account.Account, error) {
	var acc account.Account
	err := r.db.First(&acc, id).Error
	return account.Account{}, err
}

func (r repo) InsertAccount(acc account.Account) (account.Account, error) {
	err := r.db.Create(&acc).Error
	return acc, err
}
