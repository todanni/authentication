package config

import (
	"fmt"
	"testing"
)

func TestNewFromEnv(t *testing.T) {
	config, err := NewFromEnv()
	fmt.Println(config)
	fmt.Println(err)
}
