package repository

import (
	"github.com/todanni/authentication/pkg/account"
)

func (r *repo) InsertVerificationRecord(record account.VerificationRecord) (account.VerificationRecord, error) {
	err := r.db.Create(&record).Error
	return record, err
}

func (r *repo) GetVerificationRecordByCode(code string) (account.VerificationRecord, error) {
	vr := account.VerificationRecord{Code: code}
	err := r.db.Where(&vr).Last(&vr).Error
	return vr, err
}
