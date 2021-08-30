package main

import (
    "context"
    "github.com/ryanlaycock/smart-home/smart-home-api/protos/github.com/ryanlaycock/smart-home/smart-home-api/protos/sensors"
    "google.golang.org/grpc"
    "log"
)

type Reading struct {
    value int
}

func main() {
    addr := "localhost:9999"
    conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
         log.Fatal(err)
    }
    defer conn.Close()

    sensorClient := sensors.NewSensorClient(conn)
    reading, err := sensorClient.Read(context.Background(), &sensors.ReadParams{})
    if err != nil {
        log.Fatal(err)
    }

    log.Println(reading.Value)
}
