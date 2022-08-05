package config

import (
	"github.com/spf13/viper"
)

// Create mock using:
// mockgen -source=pkg/config/interface.go -destination=pkg/config/mock/mock_config.go
type Interface interface {
	Init() error
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	AllSettings() map[string]interface{}
	IsSet(key string) bool
	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetString(key string) string
	GetStringSlice(key string) []string
	UnmarshalKey(key string, rawVal interface{}, decoder ...viper.DecoderConfigOption) error
	ReadConfig(configFilePath string) error
}

const PACKAGR_SCM = "scm"
const PACKAGR_SCM_REPO_FULL_NAME = "scm_repo_full_name"
const PACKAGR_SCM_REPO_SHA = "scm_repo_sha"
const PACKAGR_SCM_REPO_TAG = "scm_repo_tag"
const PACKAGR_ENGINE_REPO_CONFIG_PATH = "engine_repo_config_path"
