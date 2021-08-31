package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    tempsensor "github.com/ryanlaycock/smart-home/smart-home-api/temp-sensor"
    "log"
    "os"
)

type Reading struct {
    value int
}

func main() {
    config, err := LoadConfig(os.Getenv("CONFIG_FILE"))
    if err != nil {
        log.Fatal(err)
        return
    }
    log.Printf("Loaded config: %+v", config)

    var tempSensors []*tempsensor.Tempsensor

    for _, tempSensor := range config.TempSensors {
        ts, err := tempsensor.New(tempSensor)
        if err != nil {
            log.Fatal(err)
            return
        }
        log.Printf("Created temperature sensor %+v", ts.Cfg.Sensors)
        tempSensors = append(tempSensors, ts)
    }

    router := gin.Default()
    router.Use(cors.Default())
    router.GET("/api/v1/sensors/temperatures", getAllTemperatures(tempSensors))

    err = router.Run()
    if err != nil {
        log.Fatal(err)
    }
}