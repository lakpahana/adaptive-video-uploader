package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/lakpahana/adaptive-video-uploader/internal/path"
)

type FFMPEG struct {
}

func (f *FFMPEG) CreateHLS(inputFilePath string, outputDirPath string) error {
	if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
		return err
	}

	command := fmt.Sprintf(createHLSCommand, inputFilePath, segmentTime, outputDirPath, path.GetFileName(inputFilePath))

	args := strings.Split(command, " ")

	commandExec := exec.Command(args[0], args[1:]...)

	_, err := commandExec.CombinedOutput()

	if err != nil {
		return fmt.Errorf("ffmpeg command failed: %s", string(err.Error()))
	}

	return nil
}

func (f *FFMPEG) CreateThumbnail(inputFilePath string, outputDirPath string) error {
	if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
		return err
	}

	command := fmt.Sprintf(createThumbnailCommand, inputFilePath, outputDirPath, path.GetFileName(inputFilePath))

	args := strings.Split(command, " ")

	commandExec := exec.Command(args[0], args[1:]...)

	_, err := commandExec.CombinedOutput()

	if err != nil {
		return fmt.Errorf("ffmpeg command failed: %s", string(err.Error()))
	}

	return nil
}
