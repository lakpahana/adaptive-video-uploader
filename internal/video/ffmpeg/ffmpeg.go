package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type FFMPEG struct {
}

func (f *FFMPEG) CreateHLS(inputFilePath string, outputDirPath string) error {
	if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
		return err
	}

	command := fmt.Sprintf(createHLSCommand, inputFilePath, segmentTime, outputDirPath)

	args := strings.Split(command, " ")

	commandExec := exec.Command(args[0], args[1:]...)

	output, err := commandExec.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("ffmpeg command failed: %s", string(err.Error()))
	}

	return nil
}

func (f *FFMPEG) CreateThumbnail(inputFilePath string, outputDirPath string) error {
	if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
		return err
	}

	command := fmt.Sprintf(createThumbnailCommand, inputFilePath, outputDirPath)

	args := strings.Split(command, " ")

	commandExec := exec.Command(args[0], args[1:]...)

	output, err := commandExec.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("ffmpeg command failed: %s", string(err.Error()))
	}

	return nil
}
