package destbuffer

import (
	"buffer-handler/logging"
	"buffer-handler/models"
)

type SftpBufferConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	RemotePath string
	LocalPath  string
}

func (cfg SftpBufferConfig) Notify(entries []models.EntryData, convertToJSONL bool) error {
	logging.Logger.Debug("Sftp Buffering")

	return nil
}
