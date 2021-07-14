package main

import (
	"log"
	"strconv"

	"github.com/wxbsocial/common"
	v1 "github.com/wxbsocial/idgenerator/adapters/api/grpc/v1"
	"github.com/wxbsocial/idgenerator/models"
)

const (
	GRPC_ADDRESS        = "GRPC_ADDRESS"
	DATA_CENTER_ID      = "DATA_CENTER_ID"
	NODE_ID             = "NODE_ID"
	BITS_DATA_CENTER_ID = "BITS_DATA_CENTER_ID"
	BITS_NODE_ID        = "BITS_NODE_ID"
	BITS_SEQUENCE       = "BITS_SEQUENCE"
	BASE_TIMESTAMP      = "BASE_TIMESTAMP"
	BUFFER_SIZE         = "BUFFER_SIZE"
)

var (
	grpc_address        = common.Getenv(GRPC_ADDRESS, "127.0.0.1:80")
	data_center_id      = common.Getenv(DATA_CENTER_ID, "0")
	node_id             = common.Getenv(NODE_ID, "0")
	bits_data_center_id = common.Getenv(BITS_DATA_CENTER_ID, "5")
	bits_node_id        = common.Getenv(BITS_NODE_ID, "5")
	bits_sequence       = common.Getenv(BITS_SEQUENCE, "12")
	base_timestamp      = common.Getenv(BASE_TIMESTAMP, "1622476800000")
	buffer_size         = common.Getenv(BUFFER_SIZE, "100")
)

func parseInt64(value string) int64 {
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func parseInt(value string) int {
	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(result)
}

func startGRPCServer() {

	config := models.NewConfig(
		parseInt64(data_center_id),
		parseInt64(node_id),
		parseInt64(bits_data_center_id),
		parseInt64(bits_node_id),
		parseInt64(bits_sequence),
		parseInt64(base_timestamp),
		parseInt(buffer_size),
	)

	generator := models.NewGenerator(config, &models.UnixTimeGetter{})

	grpcServer := NewGRPCServer(
		grpc_address,
		v1.NewIdGeneratorServer(generator))

	if err := grpcServer.Start(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	startGRPCServer()
}
