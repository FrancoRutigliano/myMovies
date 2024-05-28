package config

import (
	"errors"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT       string `mapstructure:"PORT"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (config EnvVars, err error) {
	// determinamos el modo dependiendo del valor de la variable de entorno GO_ENV
	mode := viper.GetString("GO_ENV")
	if mode == "" {
		mode = "development"
	}

	// Construimos el nombre del archivo de configuración
	envFileName := "." + mode + ".env"

	// Agregamos la ruta del directorio de configuración
	viper.AddConfigPath(".")
	// Establecemos el nombre del archivo de configuración
	viper.SetConfigName(envFileName)
	// Establecemos el tipo de archivo de configuración
	viper.SetConfigType("env")

	// Habilitamos la lectura automática de variables de entorno
	viper.AutomaticEnv()

	// Intentamos leer las configuraciones
	if err := viper.ReadInConfig(); err != nil {
		return EnvVars{}, err
	}

	// Deserializamos las configuraciones en la estructura de EnvVars
	err = viper.Unmarshal(&config)
	if err != nil {
		return EnvVars{}, err
	}

	// Realizamos validaciones de las variables de entorno
	if config.PORT == "" {
		return EnvVars{}, errors.New("PORT is not set in .env file")
	}

	if config.JWT_SECRET == "" {
		return EnvVars{}, errors.New("JWT_SECRET is not set in .env file")
	}

	return config, nil
}
