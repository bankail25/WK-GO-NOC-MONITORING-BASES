package config

import "fmt"

import (
	"github.com/spf13/viper"
)

type Config struct {
	UserNameRemedy  string `mapstructure:"USERNAME_REMEDY"`
	PasswordRemedy  string `mapstructure:"PASSWORD_REMEDY"`
	URLGetRemedy    string `mapstructure:"URL_GET_REMEDY"`
	URLUpdateRemedy string `mapstructure:"URL_UPDATE_REMEDY"`
	MongoUri        string `mapstructure:"MONGO_URI"`
}

func LoadConfig(path string) (config Config, err error) {
	// Si se proporciona un path, intentar leer archivo .env
	if path != "" {
		viper.AddConfigPath(path)
		viper.SetConfigName(".env")
		viper.SetConfigType("env")

		// Intentar leer archivo, pero no fallar si no existe
		if err = viper.ReadInConfig(); err != nil {
			// Solo mostrar mensaje si es por archivo no encontrado
			fmt.Printf("Config file not found at %s, using environment variables: %v\n", path, err)
		}
	}

	// Siempre habilitar lectura de variables de entorno
	viper.AutomaticEnv()

	// Mapear las variables de entorno a la estructura
	viper.BindEnv("USERNAME_REMEDY")
	viper.BindEnv("PASSWORD_REMEDY")
	viper.BindEnv("URL_GET_REMEDY")
	viper.BindEnv("URL_UPDATE_REMEDY")
	viper.BindEnv("MONGO_URI")

	// Unmarshal desde variables de entorno (y archivo si existe)
	err = viper.Unmarshal(&config)
	return
}
