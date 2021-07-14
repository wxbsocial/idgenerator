package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	v1pb "github.com/wxbsocial/idgenerator/adapters/api/grpc/protos/v1"
)

func TestMain(m *testing.M) {

	go startGRPCServer()
	initClient()
	defer fields.conn.Close()

	code := m.Run()

	os.Exit(code)
}

type Fields struct {
	conn   *grpc.ClientConn
	client v1pb.IdGeneratorClient
}

var fields Fields

func initClient() {

	conn, err := grpc.Dial(grpc_address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fields.conn = conn

	fields.client = v1pb.NewIdGeneratorClient(fields.conn)

}

func TestGet(t *testing.T) {

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	for i := 0; i < 36; i++ {
		result, err := fields.client.Get(ctx, &emptypb.Empty{})

		if err != nil {
			t.Fatalf("get falied %v", err)
		}

		t.Logf("get : %02X", result.Value)
	}

}
