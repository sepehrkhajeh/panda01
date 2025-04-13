package database

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	URI               string
	DBName            string `yaml:"db_name"`
	ConnectionTimeout int16  `yaml:"connection_timeout_in_ms"`
	QueryTimeout      int16  `yaml:"query_timeout_in_ms"`
}

func Load(path string) *Config {
	wd, _ := os.Getwd()
	fullPath := filepath.Join(wd, path)

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("خطا در خواندن فایل کانفیگ: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("خطا در unmarshal کردن فایل yaml: %v", err)
	}
	return &cfg
}
