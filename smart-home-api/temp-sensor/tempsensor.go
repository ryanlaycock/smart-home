package tempsensor

import (
    "fmt"
    "github.com/ryanlaycock/smart-home/smart-home-api/protos/github.com/ryanlaycock/smart-home/smart-home-api/protos/sensors"
    "google.golang.org/grpc"
)

type Cfg struct {
    Sensors        map[string]Sensor `json:"sensors"`
    ServerIP       string `json:"server_ip"`
    ServerGRPCPort string `json:"server_grpc_port"`
}

type Sensor struct {
    SensorId    string `json:"sensor_id"`
    DisplayName string `json:"display_name"`
}

type Tempsensor struct {
    conn *grpc.ClientConn
    c    sensors.SensorClient
    Cfg  Cfg
}

func New(cfg Cfg) (*Tempsensor, error) {
    t := new(Tempsensor)
    var err error
    addr := fmt.Sprintf("%v:%v", cfg.ServerIP, cfg.ServerGRPCPort)
    t.conn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, err
    }

    t.c = sensors.NewSensorClient(t.conn)
    t.Cfg = cfg
    return t, nil
}

func (t *Tempsensor) Close() error {
    return t.conn.Close()
}

