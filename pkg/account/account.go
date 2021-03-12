package account

import (
	"net/http"

	"gorm.io/gorm"
)

type Account struct {
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	ProfilePicture string      `json:"profile_picture"`
	JobTitle       string      `json:"job_title"`
	AuthDetails    AuthDetails `json:"-" gorm:"foreignKey:account_id"`
	gorm.Model
}

type Service interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
	Verify(w http.ResponseWriter, r *http.Request)
}

type Repository interface {
	InsertAuthDetails(details AuthDetails) (AuthDetails, error)
	UpdateAuthDetails(details AuthDetails) (AuthDetails, error)
	GetAuthDetails(userID int) (AuthDetails, error)
}
