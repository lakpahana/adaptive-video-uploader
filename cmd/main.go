package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/lakpahana/adaptive-video-uploader/internal/env"
	"github.com/lakpahana/adaptive-video-uploader/internal/path"
	"github.com/lakpahana/adaptive-video-uploader/internal/storage"
	"github.com/lakpahana/adaptive-video-uploader/internal/storage/firebase"
)

func main() {
	env.LoadEnv()

	storageBucket := os.Getenv("STORAGE_BUCKET")
	serviceAccount := os.Getenv("SERVICE_ACCOUNT")

	app, err := firebase.NewFirebase(context.Background(),
		path.GetProjectRootPath()+"/"+serviceAccount, storageBucket)
	if err != nil {
		fmt.Println(err)
		return
	}

	storageHandler := &storage.StorageHandler{
		Storage: app,
	}
	var fileData = bytes.NewBufferString("Hello, Go Interfaces!")
	err = storageHandler.Storage.Store(context.Background(), "hello.txt", fileData)

	if err != nil {
		fmt.Println(err)
		return
	}
}
