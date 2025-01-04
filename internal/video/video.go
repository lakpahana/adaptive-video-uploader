package video

type Video interface {
	CreateHLS(inputFilePath string, outputDirPath string) error
	CreateThumbnail(inputFilePath string, outputDirPath string) error
}

type VideoHandler struct {
	Video Video
}
