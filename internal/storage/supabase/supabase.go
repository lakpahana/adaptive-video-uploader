package supabase

import (
	"context"
	"io"

	sp_storage "github.com/supabase-community/storage-go"
)

type Supabase struct {
	Client        *sp_storage.Client
	ApiKey        string
	AppUrl        string
	StorageBucket string
}

func NewSupabase(apiKey, appUrl, storageBucket string) (*Supabase, error) {
	client := sp_storage.NewClient(appUrl, apiKey, nil)
	return &Supabase{
		Client:        client,
		ApiKey:        apiKey,
		AppUrl:        appUrl,
		StorageBucket: storageBucket,
	}, nil
}

func (s *Supabase) Store(ctx context.Context, key string, data io.Reader) (string, error) {
	upload, err := s.Client.UploadFile(s.StorageBucket, key, data)
	if err != nil {
		return "", err
	}

	return upload.Message, nil
}
