package repository

import "github.com/todanni/authentication/pkg/account"

func (r repo) GetAccount(id int) (account.Account, error) {
	var acc account.Account
	err := r.db.First(&acc, id).Error
	return account.Account{}, err
}
