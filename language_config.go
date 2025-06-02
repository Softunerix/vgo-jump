package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type LanguageConfig struct {
	PackageStartsWith    string   `yaml:"package_starts_with"`
	PackageEndsWith      string   `yaml:"package_ends_with"`
	RootMarker           []string `yaml:"root_marker"`
	LowercaseFirstLetter bool     `yaml:"lowercase_first_letter"`
	Separator            string   `yaml:"separator"`
	AddToRootMarker      string   `yaml:"add_to_root_marker"`
}

type Config struct {
	PHP  LanguageConfig `yaml:"php"`
	Java LanguageConfig `yaml:"java"`
	Py   LanguageConfig `yaml:"py"`
}

var loadedConfig *Config

func LoadConfig() error {
	data, err := os.ReadFile("languages.yaml")
	if err != nil {
		return err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}
	loadedConfig = &cfg
	return nil
}

func GetLangKey(path string) string {
	ext := strings.ToLower(path[strings.LastIndex(path, ".")+1:])
	switch ext {
	case "php":
		return "php"
	case "java":
		return "java"
	case "py":
		return "py"
	default:
		return ""
	}
}

func GetConfigFor(path string) (*LanguageConfig, string) {
	if loadedConfig == nil {
		return nil, ""
	}

	switch GetLangKey(path) {
	case "php":
		return &loadedConfig.PHP, "php"
	case "java":
		return &loadedConfig.Java, "java"
	case "py":
		return &loadedConfig.Py, "py"
	default:
		return nil, ""
	}
}
