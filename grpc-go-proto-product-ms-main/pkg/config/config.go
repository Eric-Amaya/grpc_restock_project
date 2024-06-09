package config

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
    _ "github.com/lib/pq"
)

type Config struct {
    Port   string `mapstructure:"PORT"`
    DBHost string `mapstructure:"DB_HOST"`
    DBPort string `mapstructure:"DB_PORT"`
    DBUser string `mapstructure:"DB_USER"`
    DBPass string `mapstructure:"DB_PASS"`
    DBName string `mapstructure:"DB_NAME"`
}

func LoadConfig() (Config, error) {
    var config Config

    // Indicar a Viper que use el formato .env
    viper.SetConfigFile(".env")
    
    // Leer las variables de entorno del archivo .env
    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading .env file: %v", err)
    }

    // Habilitar la lectura de variables de entorno del sistema
    viper.AutomaticEnv()

    // Unmarshal las variables de entorno en la estructura Config
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
        return config, err
    }
    fmt.Println("Variables de entorno cargadas:", viper.AllSettings())

    return config, nil
}
