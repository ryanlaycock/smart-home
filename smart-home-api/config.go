package main

import (
    "encoding/json"
    tempsensor "github.com/ryanlaycock/smart-home/smart-home-api/temp-sensor"
    "os"
)

type Config struct {
    TempSensors []tempsensor.Cfg `json:"temp_sensors"`
}

func LoadConfig(file string) (Config, error) {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        return Config{}, err
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        return Config{}, err
    }
    return config, nil
}