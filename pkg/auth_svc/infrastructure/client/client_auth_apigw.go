package client_apigw_auth

import (
	"context"
	"fmt"

	"github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/pb"
	config_apigw "github.com/shaan/socialMediaApiGateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitAuthClient(config *config_apigw.Config) (*pb.AuthServiceClient, error) {
	fmt.Println("AuthSvcUrl from config:", config.AuthSvcUrl)

	cc, err := grpc.NewClient(config.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("----", err)
		return nil, err
	}

	Client := pb.NewAuthServiceClient(cc)
	// Run gRPC Health Check
	healthClient := grpc_health_v1.NewHealthClient(cc)

	resp, err := healthClient.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{
		Service: "auth_proto.AuthService",
	})
	if err != nil {
		fmt.Println("gRPC Health Check failed:", err)
	} else {
		fmt.Println("✅ AuthService health:", resp.Status.String())
	}
	// 	resp, err := Client.Ping(context.Background(), &pb.PingRequest{})
	// if err != nil {
	//     fmt.Println("❌ gRPC Ping failed:", err)
	// } else {
	//     fmt.Println("✅ gRPC Ping success:", resp.Message)
	// }

	return &Client, nil
}
