package server

import (
	"fmt"
	"time"
)

type Settings struct {
	Port    int           `json:"port"`
	Timeout time.Duration `json:"timeout"`
}

func DefaultSettings() *Settings {
	return &Settings{
		Port:    8080,
		Timeout: 60,
	}
}

func (s *Settings) Validate() error {
	if s.Port == 0 {
		return fmt.Errorf("validate Settings: Port missing")
	}
	return nil
}
