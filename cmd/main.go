package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lakpahana/adaptive-video-uploader/internal/env"
	"github.com/lakpahana/adaptive-video-uploader/internal/path"
	"github.com/lakpahana/adaptive-video-uploader/internal/storage"
	"github.com/lakpahana/adaptive-video-uploader/internal/storage/firebase"
	"github.com/lakpahana/adaptive-video-uploader/internal/video"
	"github.com/lakpahana/adaptive-video-uploader/internal/video/ffmpeg"
)

func main() {
	env.LoadEnv()

	ffmpeg := &ffmpeg.FFMPEG{}

	video := &video.VideoHandler{
		Video: ffmpeg,
	}

	inputVideoPath := path.GetProjectRootPath() + "/cmd/input.mp4"
	outputDirPath := path.GetProjectRootPath() + "/cmd/output"

	err := video.Video.CreateHLS(inputVideoPath, outputDirPath)

	if err != nil {
		fmt.Println(err)
	}

	err = video.Video.CreateThumbnail(inputVideoPath, outputDirPath)

	if err != nil {
		fmt.Println(err)
	}

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

	for _, file := range path.GetFiles(outputDirPath) {
		data, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		fileName := path.GetFileName(file)

		link, err := storageHandler.Storage.Store(context.Background(), fileName, data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(link)
	}
}
