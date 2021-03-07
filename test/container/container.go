package container

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

type Container struct {
	container testcontainers.Container
}

func (c *Container) Close(ctx context.Context) error {
	return c.container.Terminate(ctx)
}
