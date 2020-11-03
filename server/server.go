package main

import (
	"go-grpc-stream/server/sensorpb"
	"net"

	"google.golang.org/grpc"
	"github.com/sirupsen/logrus"
)

type server struct{}

func (*server) TempSensor(req *sensorpb.SensorRequest,
		stream sensorpb.Sensor_TempSensorServer) error {

			return nil;
}

func (*server) HumiditySensor(req *sensorpb.SensorRequest,
		stream sensorpb.Sensor_TempSensorServer) error {

			return nil;


}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.8080")
	
	if err != nil {
		log.Fatal("Error while listening : %v", err)	
	}
	
	s := grpc.NewServer(
		sensorpb.RegisterSensorServer(s, &server{}
	)
	
	if err s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}