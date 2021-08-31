package main

//go:generate protoc -I=../ --go-grpc_out=protos/ --go_out=protos/ ../protos/sensors.proto
