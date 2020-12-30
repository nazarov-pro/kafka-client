package conf

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config app configs
var Config *viper.Viper

// LoadConfiguration loads env vars and other configurations
func LoadConfiguration() {
	Config = viper.New()
	confFile := os.Getenv("CONFIG_FILE")
	if confFile != "ENV_ONLY" {
		if confFile == "" {
			confFile = `configs/app/config.yaml`
		}
		Config.SetConfigFile(confFile)
		err := Config.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	Config.SetEnvPrefix("_")
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	Config.AutomaticEnv()
}
