package utils

import(
	"github.com/spf13/viper"
)

type Config struct {
    ManagerAccessApiUrl      string `mapstructure:"MANAGEMENT_SERVER_API_URL"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)

	if err != nil {
		return 
	}
	
    return
}

