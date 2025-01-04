package storage

import (
	"context"
	"io"
)

type Storage interface {
	Store(ctx context.Context, key string, data io.Reader) error
}

type StorageHandler struct {
	Storage Storage
}
