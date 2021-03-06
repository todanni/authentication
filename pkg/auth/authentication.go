package auth

import (
	"net/http"
)

type AuthenticationDetails struct {
	AccountID uint   `json:"userID" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"not null,unique"`
	Password  string `json:"password" gorm:"not null"`
	Verified  bool   `json:"verified" gorm:"default:false"`
}

type Service interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

type Repository interface {
	Insert(details AuthenticationDetails) (AuthenticationDetails, error)
	Update(details AuthenticationDetails) (AuthenticationDetails, error)
	Get(userID int) (AuthenticationDetails, error)
}
