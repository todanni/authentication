package repository

import (
	"github.com/todanni/authentication/pkg/account"
)

func (r *repo) InsertVerificationRecord(record account.VerificationRecord) (account.VerificationRecord, error) {
	err := r.db.Create(&record).Error
	return record, err
}

func (r *repo) GetVerificationRecord(accountID int) (account.VerificationRecord, error) {
	panic("implement me")
}

func (r *repo) UpdateVerificationRecord(code string) (account.VerificationRecord, error) {
	// UPDATE auth_details where account_id = ? SET verified = true



	//vr := account.VerificationRecord{
	//	AccountID: 0,
	//	Code:      "",
	//	Model:     gorm.Model{},
	//}
	//err := r.db.Where(account.VerificationRecord{Code: code}).Update("verified", true).Error
	//return verificationRecord, err
	panic("")
}

