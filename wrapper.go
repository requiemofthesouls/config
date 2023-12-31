package config

import (
	"strings"

	"github.com/spf13/viper"
)

func New(cfgPath string) (Wrapper, error) {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvPrefix("ENV")
	v.SetEnvKeyReplacer(
		strings.NewReplacer(".", "_"),
	)
	v.SetConfigFile(cfgPath)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return &wrapper{
		viper: v,
	}, nil
}

type (
	Wrapper interface {
		GetBool(key string) bool
		GetString(key string) string
		GetInt64(key string) int64
		IsSet(key string) bool
		UnmarshalKey(key string, rawVal interface{}) error
	}

	wrapper struct {
		viper *viper.Viper
	}
)

func (w *wrapper) GetBool(key string) bool {
	return w.viper.GetBool(key)
}

func (w *wrapper) GetString(key string) string {
	return w.viper.GetString(key)
}

func (w *wrapper) GetInt64(key string) int64 {
	return w.viper.GetInt64(key)
}

func (w *wrapper) IsSet(key string) bool {
	return w.viper.IsSet(key)
}

func (w *wrapper) UnmarshalKey(key string, rawVal interface{}) error {
	return w.viper.UnmarshalKey(key, rawVal)
}
