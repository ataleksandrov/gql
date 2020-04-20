package storage

import "fmt"

type Settings struct {
	Type              string `mapstructure:"type" description:"Type of the storage"`
	URI               string `mapstructure:"uri" description:"URI of the storage"`
	SkipSSLValidation bool   `mapstructure:"skip_ssl_validation" description:"whether to skip ssl verification when connecting to the storage"`
}

func DefaultSettings() *Settings {
	return &Settings{
		Type:              "",
		URI:               "",
		SkipSSLValidation: false,
	}
}

func (s *Settings) Validate() error {
	if len(s.Type) == 0 {
		return fmt.Errorf("storage type must be provided")
	}
	if len(s.URI) == 0 {
		return fmt.Errorf("storage URI must be provided")
	}
	return nil
}
