package supabase

import (
	"context"
	"io"
	"strings"

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

func (s *Supabase) Store(ctx context.Context, key string, folderName string, data io.Reader) (string, error) {

	contentType := "text/plain;charset=UTF-8"

	if strings.Contains(key, ".png") {
		contentType = "image/png"
	}

	if strings.Contains(key, ".jpg") {
		contentType = "image/jpeg"
	}

	_, err := s.Client.UploadFile(s.StorageBucket+"/"+folderName, key, data, sp_storage.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		return "", err
	}
	uploadURL := s.AppUrl + "/object/public/" + s.StorageBucket + "/" + folderName + "/" + key
	return uploadURL, nil
}
