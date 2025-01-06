package storage

import (
	"context"
	"io"
)

type Storage interface {
	Store(ctx context.Context, key string, folderName string, data io.Reader) (string, error)
}

type StorageHandler struct {
	Storage Storage
}
