package env

import (
	"fmt"

	"github.com/ataleksand/gql/storage"

	"github.com/ataleksand/gql/server"
	"github.com/spf13/viper"
)

type ConfigFile struct {
	Name     string `description:"name of the configuration file"`
	Location string `description:"location of the configuration file"`
	Format   string `description:"extension of the configuration file"`
}

type Env struct {
	Storage *storage.Settings
	Server  *server.Settings
}

func (s *Env) Validate() error {
	validatable := []interface {
		Validate() error
	}{s.Server, s.Storage}
	for _, v := range validatable {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func DefaultSettings() *Env {
	return &Env{
		Server:  server.DefaultSettings(),
		Storage: storage.DefaultSettings(),
	}
}

func DefaultConfigFile() ConfigFile {
	return ConfigFile{
		Name:     "application",
		Location: ".",
		Format:   "yml",
	}
}

func New() (*Env, error) {
	config := DefaultSettings()
	configFile := DefaultConfigFile()

	v := viper.New()

	v.AddConfigPath(configFile.Location)
	v.SetConfigName(configFile.Name)
	v.SetConfigType(configFile.Format)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("could not read configuration cfg: %s", err)
		}
	}

	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error loading configuration: %s", err)
	}

	return config, nil
}
