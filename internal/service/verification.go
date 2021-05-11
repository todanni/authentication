package service

import (
	"github.com/todanni/alerts"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/authentication/pkg/account"
)

func (s *service) Verify(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if code == "" {
		http.Error(w, "missing verification code", http.StatusBadRequest)
		return
	}

	var vr account.VerificationRecord
	vr, err := s.repo.GetVerificationRecordByCode(code)
	if err != nil {
		http.Error(w, "invalid code", http.StatusNotFound)
		return
	}

	//TODO: If code is older than 24 hours, invalidate and generate new
	if vr.CreatedAt.After(vr.CreatedAt.Add(time.Hour * 24)) {
		log.Error("this code has expired. failed to verify account.")
		return
	}

	err = s.repo.SetAuthDetailsValid(vr.AccountID)
	if err != nil {
		http.Error(w, "couldn't verify account", http.StatusInternalServerError)
		return
	}
	str := strconv.FormatUint(uint64(vr.AccountID), 10)
	//TODO: Send activation alert using alerter
	err = s.alerter.SendActivationAlert(alerts.RegisterRequest{
		Email:    vr.Code,
		FullName: str,
	})

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
