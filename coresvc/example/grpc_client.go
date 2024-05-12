package main

import (
	"context"
	"fmt"
	airlinev1 "github.com/nofendian17/openota/coresvc/gen/go/proto/airline/v1"
	"github.com/nofendian17/openota/coresvc/gen/go/proto/healthcheck"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {

	addr := "localhost:3000"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	healthCheckServiceClient := healthcheck.NewHealthCheckServiceClient(conn)

	responsePing, err := healthCheckServiceClient.Ping(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Could not create request : %v", err)
	}

	fmt.Println(responsePing)

	responseHealth, err := healthCheckServiceClient.HealthCheck(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(responseHealth)

	responseReadiness, err := healthCheckServiceClient.Readiness(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(responseReadiness)

	airlineServiceClient := airlinev1.NewAirlineServiceClient(conn)
	airlines, err := airlineServiceClient.GetAll(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(airlines.GetAirlines())

	getByID, err := airlineServiceClient.GetByID(ctx, &airlinev1.GetByIDRequest{
		Id: "not-valid",
	})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}
	fmt.Println(getByID)
}
