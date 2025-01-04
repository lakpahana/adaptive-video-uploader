package firebase

import (
	"context"
	"fmt"
	"io"

	firebaselib "firebase.google.com/go"

	"google.golang.org/api/option"
)

type Firebase struct {
	app           *firebaselib.App
	storageBucket string
}

func NewFirebase(ctx context.Context, credPath string, storageBucket string) (*Firebase, error) {
	opt := option.WithCredentialsFile(credPath)
	app, err := firebaselib.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("firebase.NewApp: %v", err)
	}
	return &Firebase{app: app,
		storageBucket: storageBucket}, nil
}

func (f *Firebase) Store(ctx context.Context, key string, data io.Reader) error {
	client, err := f.app.Storage(ctx)
	if err != nil {
		return fmt.Errorf("app.Storage: %v", err)
	}

	bucket, err := client.Bucket(f.storageBucket)
	if err != nil {
		return fmt.Errorf("client.DefaultBucket: %v", err)
	}

	obj := bucket.Object(key)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, data); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("w.Close: %v", err)
	}
	return nil
}
