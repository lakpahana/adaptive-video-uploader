package path

import (
	"os"
	"regexp"
)

const projectDirName = "adaptive-video-uploader"

func GetProjectRootPath() string {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return string(rootPath)
}
