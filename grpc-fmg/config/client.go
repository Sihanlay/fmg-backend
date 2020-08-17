package config

import (
	"context"
	"google.golang.org/grpc"
	Config2 "grpc-demo/config/proto"
)

var client Config2.ConfigServiceClient

func InitConfig() {
	conn, err := grpc.Dial("106.52.218.85:5432", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = Config2.NewConfigServiceClient(conn)
}

func Config() []byte {
	c, err := client.Config(context.TODO(), &Config2.ConfigRequest{
		Mode: "dev",
	})

	if err != nil {
		panic(err)
	}
	return []byte(c.Config)
}

