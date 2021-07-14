package main

import (
	"log"
	"net"

	v1 "github.com/wxbsocial/idgenerator/adapters/api/grpc/v1"
	"google.golang.org/grpc"

	v1pb "github.com/wxbsocial/idgenerator/adapters/api/grpc/protos/v1"
)

type GRPCServer struct {
	address             string
	idGeneratorServerV1 *v1.IdGeneratorServer
}

func NewGRPCServer(address string,
	idGeneratorServer *v1.IdGeneratorServer,
) *GRPCServer {

	return &GRPCServer{
		address:             address,
		idGeneratorServerV1: idGeneratorServer,
	}
}

func (s *GRPCServer) Start() error {

	lis, err := net.Listen("tcp", s.address)

	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}

	server := grpc.NewServer()
	v1pb.RegisterIdGeneratorServer(server, s.idGeneratorServerV1)

	log.Printf("start to listen: %v", s.address)
	if err := server.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}

	return nil
}
