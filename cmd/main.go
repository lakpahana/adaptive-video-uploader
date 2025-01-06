package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"github.com/lakpahana/adaptive-video-uploader/internal/env"
	"github.com/lakpahana/adaptive-video-uploader/internal/path"
	"github.com/lakpahana/adaptive-video-uploader/internal/storage"
	"github.com/lakpahana/adaptive-video-uploader/internal/storage/supabase"
	"github.com/lakpahana/adaptive-video-uploader/internal/video"
	"github.com/lakpahana/adaptive-video-uploader/internal/video/ffmpeg"
)

func main() {
	env.LoadEnv()

	sp := &supabase.Supabase{
		ApiKey:        os.Getenv("SUPABASE_KEY"),
		AppUrl:        os.Getenv("SUPABASE_URL"),
		StorageBucket: os.Getenv("SUPABASE_BUCKET"),
	}

	app, err := supabase.NewSupabase(sp.ApiKey, sp.AppUrl, sp.StorageBucket)

	if err != nil {
		fmt.Println(err)
		return
	}

	ffmpeg := &ffmpeg.FFMPEG{}

	video := &video.VideoHandler{
		Video: ffmpeg,
	}

	inputVideoPath := path.GetProjectRootPath() + "/cmd/daff.mp4"
	outputDirPath := path.GetProjectRootPath() + "/cmd/output"

	if _, err := os.Stat(outputDirPath); !os.IsNotExist(err) {
		cmd := exec.Command("rm", "-rf", outputDirPath)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Deleted output directory")
	}

	// err = video.Video.CreateDASH(inputVideoPath, outputDirPath)

	if err != nil {
		fmt.Println(err)
	}

	err = video.Video.CreateThumbnail(inputVideoPath, outputDirPath)

	if err != nil {
		fmt.Println(err)
	}

	storageHandler := &storage.StorageHandler{
		Storage: app,
	}

	folderName := uuid.New().String()

	for _, file := range path.GetFiles(outputDirPath) {
		data, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		fileName := path.GetFileName(file)

		link, err := storageHandler.Storage.Store(context.Background(), fileName, folderName, data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(link)
	}
}

// storageBucket := os.Getenv("STORAGE_BUCKET")
// serviceAccount := os.Getenv("SERVICE_ACCOUNT")

// app, err := firebase.NewFirebase(context.Background(),
// 	path.GetProjectRootPath()+"/"+serviceAccount, storageBucket)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// ftpPort, err := strconv.Atoi(os.Getenv("FTP_PORT"))

// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// ftpConf := &ftp.FTPConf{
// 	Host:     os.Getenv("FTP_HOST"),
// 	Port:     ftpPort,
// 	Username: os.Getenv("FTP_USER"),
// 	Password: os.Getenv("FTP_PASSWORD"),
// 	Path:     os.Getenv("FTP_PATH"),
// }

// ftp, err := ftp.NewFTP(ftpConf)

// if err != nil {
// 	fmt.Println(err)
// 	return
// }
