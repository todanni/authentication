package container

import (
	"crypto/rand"
	"encoding/base32"
)

type credentials struct {
	username string
	password string
}

func newCredentials() credentials {
	randomBytes := make([]byte, 32)
	_, _ = rand.Read(randomBytes)

	pw := base32.StdEncoding.EncodeToString(randomBytes)[:20]

	return credentials{
		username: "integration_tests",
		password: pw,
	}
}
