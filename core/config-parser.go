package core

import (
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
	Net struct {
		BindIp string `yaml:"bindIp"`
		WebPort string `yaml:"webPort"`
		ApiPort string `yaml:"apiPort"`
	} `yaml:"net"`
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
