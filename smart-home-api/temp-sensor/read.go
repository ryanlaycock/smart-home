package tempsensor

import (
    "context"
    "errors"
    "github.com/ryanlaycock/smart-home/smart-home-api/protos/github.com/ryanlaycock/smart-home/smart-home-api/protos/sensors"
)

var (
    sensorNotFoundErr = errors.New("sensor not found")
)

func (t *Tempsensor) Read(sensorId string) (float64, error) {
    if _, ok := t.Cfg.Sensors[sensorId]; !ok {
        return 0, sensorNotFoundErr
    }

    reading, err := t.c.ReadTemp(context.Background(), &sensors.TempSensor{SensorId: t.Cfg.Sensors[sensorId].SensorId})
    if err != nil {
        return 0, err
    }
    
    return reading.Value, nil
}