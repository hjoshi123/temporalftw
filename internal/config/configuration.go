package config

import "github.com/spf13/viper"

const (
	DevelopmentEnvironment = "development"
	ProductionEnvironment  = "production"
)

type Configuration struct {
	DBHost      string `mapstructure:"db_host"`
	DBPort      string `mapstructure:"db_port"`
	DBUser      string `mapstructure:"db_user"`
	DBName      string `mapstructure:"db_name"`
	DBPassword  string `mapstructure:"db_password"`
	LogLevel    string `mapstructure:"log_level"`
	Environment string `mapstructure:"environment"`
}

var Spec *Configuration

func IsDevelopment() bool {
	return Spec.Environment == DevelopmentEnvironment
}

func init() {
	Spec = new(Configuration)

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Spec)
}
