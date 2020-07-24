package core

import (
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
	Net struct {
		WebPort string `yaml:"webPort"`
		ApiHost string `yaml:"apiHost"`
		ApiPort string `yaml:"apiPort"`
	} `yaml:"net"`

	Auth []string `yaml:"auth"`

	Log struct {
		MaxFiles int `yaml:"maxFiles"`
  	MaxFileSize int `yaml:"maxFileSize"`
  	RetentionDays int `yaml:"retentionDays"`
	} `yaml:"log"`
}

func GetConfig(configPath string) (*Config, error) {
    // Create config structure
    config := &Config{}

    // Open config file
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Init new YAML decode
    d := yaml.NewDecoder(file)

    // Start YAML decoding from file
    if err := d.Decode(&config); err != nil {
        return nil, err
    }

    return config, nil
}
