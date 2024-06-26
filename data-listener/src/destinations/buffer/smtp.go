package buffer

import (
	"datalistener/src/logging"
	"datalistener/src/models"
)

type SmtpsBufferConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	UseTLS   bool
}

func (cfg *SmtpsBufferConfig) Notify(entries []models.EntryData, convertToJSONL bool) error {
	logging.Logger.Debug("Smtps Buffering")

	return nil
}
