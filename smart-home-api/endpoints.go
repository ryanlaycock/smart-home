package main

import (
    "github.com/gin-gonic/gin"
    tempsensor "github.com/ryanlaycock/smart-home/smart-home-api/temp-sensor"
    "log"
    "net/http"
)

type TempSensorResp struct {
    DisplayName string `json:"display_name"`
    ID string `json:"id"`
    RawValue float64 `json:"raw_value"`
}

func getAllTemperatures(tempSensors []*tempsensor.Tempsensor) gin.HandlerFunc {
    return func (c *gin.Context) {
        var resp []TempSensorResp

        for _, ts := range tempSensors {
            for k, v := range ts.Cfg.Sensors {
                log.Printf("Reading sensor %v", v.SensorId)
                value, err := ts.Read(k)
                if err != nil {
                    log.Println(err)
                    continue
                }
                resp = append(resp, TempSensorResp{
                    DisplayName: v.DisplayName,
                    ID:          k,
                    RawValue:    value,
                })
            }
        }

        if len(resp) == 0 {
            c.JSON(http.StatusNotFound, nil)
            return
        }
        c.JSON(http.StatusOK, resp)
    }
}
