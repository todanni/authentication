package container

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PgContainer struct {
	Container
	credentials
	db string
}

const (
	timeoutDeadline = time.Minute * 3
)

// NewPGContainer creates a new container running PG and blocks until it is ready to use.
func NewPGContainer(db string) (*PgContainer, error) {
	ctx := context.Background()
	credentials := newCredentials()
	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"54325:5432"},
		WaitingFor: wait.ForSQL("5432", "postgres", func(port nat.Port) string {
			return fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
				port.Int(), credentials.username, credentials.password, db)
		}).Timeout(timeoutDeadline),
		Env: map[string]string{
			"POSTGRES_DB":       db,
			"POSTGRES_USER":     credentials.username,
			"POSTGRES_PASSWORD": credentials.password,
		},
	}

	pg, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot start PG test container: %v", err)
	}

	return &PgContainer{Container{pg}, credentials, db}, nil
}

func (p *PgContainer) ConnectionString() (string, error) {
	endpoint, err := p.container.Endpoint(context.Background(), "")
	if err != nil {
		return "", fmt.Errorf("cannot get endpoint: %v", err)
	}

	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", p.username, p.password, endpoint, p.db), nil
}
