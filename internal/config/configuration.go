package config

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

const (
	DevelopmentEnvironment = "development"
	ProductionEnvironment  = "production"
)

type Configuration struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBName           string `mapstructure:"DB_NAME"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	LogLevel         string `mapstructure:"LOG_LEVEL"`
	Environment      string `mapstructure:"ENVIRONMENT"`
	Port             int    `mapstructure:"PORT"`
	TemporalHostPort string `mapstructure:"TEMPORAL_HOST_PORT"`
}

var Spec *Configuration

func IsDevelopment() bool {
	return Spec.Environment == DevelopmentEnvironment
}

func init() {
	Spec = new(Configuration)

	v := viper.New()

	v.AutomaticEnv()
	v.AddConfigPath(".")
	v.SetConfigType("env")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	bindEnvs(v, Spec)

	err = v.Unmarshal(&Spec)
}

func bindEnvs(vip *viper.Viper, iFace interface{}) {
	ifv := reflect.ValueOf(iFace)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}

	for i := 0; i < ifv.NumField(); i++ {
		v := ifv.Field(i)
		t := ifv.Type().Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}

		switch v.Kind() {
		default:
			vip.BindEnv(tv, tv)
		}
	}
}
