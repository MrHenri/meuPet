package configs

import (
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	WebServerPort  string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort string `mapstructure:"GRPC_SERVER_PORT"`
}

func Load() (*conf, error) {

	// Get the file path of the current function call
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get module path")
	}

	// Remove the file name from the path
	dirPath := path.Dir(filePath)

	var cfg *conf
	viper.SetConfigFile(dirPath + "/.env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	return cfg, err
}
