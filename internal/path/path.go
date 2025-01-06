package path

import (
	"os"
	"path/filepath"
	"regexp"
)

const projectDirName = "adaptive-video-uploader"

func GetProjectRootPath() string {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return string(rootPath)
}

func GetFiles(dirPath string) []string {
	var files []string
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func GetFileName(filePath string) string {
	_, fileName := filepath.Split(filePath)
	return fileName
}
