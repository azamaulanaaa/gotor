package main

import (
    "strings"
    "github.com/spf13/viper"
)

type Config struct {
    HttpServe       HttpServeConfig
    TorrentClient   TorrentClientConfig
}

type HttpServeConfig struct {
    Port        uint16
    Timeout     uint16
}

type TorrentClientConfig struct {
    PieceLifetime           uint32
    CleanUpPeaceInterval    uint32
    MemoryStorage           bool
    StorageDir              string
}

func DefaultConfig() Config {
    return Config{
        HttpServe: HttpServeConfig{
            Port: 8080,
            Timeout: 30,
        },
        TorrentClient: TorrentClientConfig{
            PieceLifetime:          10,
            CleanUpPeaceInterval:   1,
            MemoryStorage:          false,
            StorageDir:             "./cache",
        },
    }
}

func LoadConfig() Config {
    config := DefaultConfig()
    viper_config := viper.New()

    viper_config.SetConfigName("config")
    viper_config.SetConfigType("json")
    viper_config.AddConfigPath(".")
    viper_config.SetEnvPrefix("gotor")
    viper_config.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

    viper_config.ReadInConfig()
    viper_config.AutomaticEnv()

    viper_config.Unmarshal(&config)
    return config
}
